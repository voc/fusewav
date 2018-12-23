package fs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/voc/fusewav/wav"
)

// A WavFs represents a File-System serving the assembled Wav-Files
type WavFs struct {
	server *fuse.Server
}

// NewWavFs constructs a new WavFs-Instance for a given Wav-Reader
func NewWavFs(reader *wav.Reader, mountpoint string) (*WavFs, error) {
	impl := &WavFsImpl{
		FileSystem: pathfs.NewDefaultFileSystem(),
		reader:     reader,
	}

	pathNodeFs := pathfs.NewPathNodeFs(impl, nil)
	server, _, err := nodefs.MountRoot(mountpoint, pathNodeFs.Root(), nil)

	if err != nil {
		return nil, err
	}

	fs := &WavFs{server: server}
	return fs, nil
}

// Serve starts serving the WavFs on the given mountpoint
func (fs *WavFs) Serve() {
	fs.server.Serve()
}
