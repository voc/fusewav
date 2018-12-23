package wav

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

const headerSize = int64(44)
const offsetFileSize = int64(4)
const offsetDataBlockSize = int64(40)

// The AggregatedWavFile Object represents a set of aggregated Wav-Files in the given directory
type AggregatedWavFile struct {
	files     []os.FileInfo
	base      string
	directory string
}

// NewAggregatedWavFile constructs a new AggregatedWavFile Object
func NewAggregatedWavFile(base string, directory string, files []os.FileInfo) *AggregatedWavFile {
	file := &AggregatedWavFile{
		files:     files,
		base:      base,
		directory: directory,
	}
	return file
}

// Filesize calculates the exact size of a Wav-Header plus all Data-Bytes from all aggregated Wav-Files
func (f *AggregatedWavFile) Filesize() uint64 {
	var sum uint64
	for index, file := range f.files {
		if index == 0 {
			sum += uint64(file.Size())
		} else {
			sum += uint64(file.Size() - headerSize)
		}
	}
	return sum
}

// Read performs a read in the reconstructed Header or the Wav-File(s) backing the requested Area
func (f *AggregatedWavFile) Read(buf []byte, off int64) error {
	inOffset := off
	outOffset := int64(0)
	outLen := int64(len(buf))

	fmt.Printf("request to read %d bytes from offset %d\n", outLen, inOffset)

	if inOffset < headerSize {
		header := f.generateHeader()
		copyLen := min(headerSize-inOffset, outLen)

		fmt.Printf("part of the request overlaps with the header. "+
			"copying %d bytes from the header starting at offset %d "+
			"to the output-buffer starting at offset %d\n",
			copyLen, inOffset, outOffset)

		copied := copy(buf[outOffset:], header[inOffset:(inOffset+copyLen)])
		if int64(copied) != copyLen {
			return fmt.Errorf("Only copied %d byted from header instead of the expected %d bytes", copied, copyLen)
		}

		inOffset = 0
		outOffset += copyLen
		outLen -= copyLen
	} else {
		inOffset -= headerSize
	}

	for _, file := range f.files {
		fileSize := file.Size() - headerSize

		if inOffset < fileSize {
			copyLen := min(fileSize-inOffset, outLen)
			if copyLen <= 0 {
				// nothing to copy
				break
			}

			inOffsetPlusHeader := inOffset + headerSize
			filebuf := f.readFile(file.Name(), inOffsetPlusHeader, copyLen)
			copied := copy(buf[outOffset:], filebuf)

			fmt.Printf("copying from %s: %d bytes starting at offset %d "+
				"to the output-buffer starting at offset %d (file-length is %d)\n",
				file.Name(), copyLen, inOffsetPlusHeader, outOffset, file.Size())

			if int64(copied) != copyLen {
				return fmt.Errorf("Only copied %d byted from header instead of the expected %d bytes", copied, copyLen)
			}

			inOffset = 0
			outOffset += copyLen
			outLen -= copyLen
		} else {
			fmt.Printf("skipping %s because the requested Offset %d is outside of this files data-length %d\n",
				file.Name(), inOffset, fileSize)

			inOffset -= fileSize
		}
	}

	if outLen != 0 {
		return fmt.Errorf("could not completely satisfy read-request, %d bytes remain ZERO", outLen)
	}

	return nil
}

func (f *AggregatedWavFile) generateHeader() []byte {
	filepath := filepath.Join(f.base, f.directory, f.files[0].Name())
	buf := make([]byte, headerSize)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Unable to open File %s: %s", filepath, err)
		return buf
	}

	file.Read(buf)

	binary.LittleEndian.PutUint32(buf[offsetFileSize:], uint32(f.Filesize()-8))
	binary.LittleEndian.PutUint32(buf[offsetDataBlockSize:], uint32(f.Filesize()-44))

	return buf
}

func (f *AggregatedWavFile) readFile(filename string, offset int64, length int64) []byte {
	filepath := filepath.Join(f.base, f.directory, filename)
	buf := make([]byte, length)

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Unable to open File %s: %s", filepath, err)
		return buf
	}
	file.Seek(offset, 0)
	file.Read(buf)
	return buf
}

// max returns the larger of x or y.
func min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}
