package wav

import (
	"os"
)

type AggregatedWavFile struct {
	files []os.FileInfo
}

func NewAggregatedWavFile(files []os.FileInfo) *AggregatedWavFile {

	file := &AggregatedWavFile{
		files: files,
	}
	return file
}

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

func (f *AggregatedWavFile) Read(buf []byte, off int64) {

}
