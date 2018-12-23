package fs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

func wavFileRead(fs *WavFsImpl, _ []string) (file nodefs.File, code fuse.Status) {
	return nodefs.NewDataFile([]byte("")), fuse.OK
}

func wavFileAttr(_ *WavFsImpl, _ []string) (*fuse.Attr, fuse.Status) {
	return &fuse.Attr{
		Mode: fuse.S_IFREG | 0644,
		Size: 0,
	}, fuse.OK
}
