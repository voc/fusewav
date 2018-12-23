package fs

import (
	"fmt"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/voc/fusewav/wav"
)

func wavFileRead(fs *WavFsImpl, matches []string) (file nodefs.File, code fuse.Status) {
	directory := findDirectoryForFilename(fs.reader, matches[1])
	if directory == "" {
		return nil, fuse.ENOENT
	}

	aggregatedWavFile, err := fs.reader.GetAggregatedWavFile(directory)
	if err != nil {
		fmt.Printf("Error creating aggregated File: %s", err)
		return nil, fuse.ENOENT
	}

	return newDelegatingFuseFile(aggregatedWavFile), fuse.OK
}

func wavFileAttr(fs *WavFsImpl, matches []string) (*fuse.Attr, fuse.Status) {
	directory := findDirectoryForFilename(fs.reader, matches[1])
	if directory == "" {
		return nil, fuse.ENOENT
	}

	aggregatedWavFile, err := fs.reader.GetAggregatedWavFile(directory)
	if err != nil {
		fmt.Printf("Error creating aggregated File: %s", err)
		return nil, fuse.EIO
	}

	return &fuse.Attr{
		Mode: fuse.S_IFREG | 0644,
		Size: aggregatedWavFile.Filesize(),
	}, fuse.OK
}

type delegatingFuseFile struct {
	aggregatedWavFile *wav.AggregatedWavFile

	nodefs.File
}

func newDelegatingFuseFile(aggregatedWavFile *wav.AggregatedWavFile) nodefs.File {
	f := new(delegatingFuseFile)
	f.aggregatedWavFile = aggregatedWavFile
	f.File = nodefs.NewDefaultFile()
	return f
}

func (f *delegatingFuseFile) GetAttr(out *fuse.Attr) fuse.Status {
	out.Mode = fuse.S_IFREG | 0644
	out.Size = f.aggregatedWavFile.Filesize()
	return fuse.OK
}

func (f *delegatingFuseFile) Read(buf []byte, off int64) (res fuse.ReadResult, code fuse.Status) {
	err := f.aggregatedWavFile.Read(buf, off)
	if err != nil {
		return nil, fuse.EIO
	}

	return fuse.ReadResultData(buf), fuse.OK
}
