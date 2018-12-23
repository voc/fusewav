package fs

import (
	"strings"

	"github.com/hanwen/go-fuse/fuse"
)

func wavDirList(fs *WavFsImpl, _ []string) ([]fuse.DirEntry, fuse.Status) {
	mathingDirectories := fs.reader.ListMatchingDirectories()
	entries := make([]fuse.DirEntry, len(mathingDirectories))

	for i, directory := range mathingDirectories {
		entries[i] = fuse.DirEntry{
			Name: strings.Replace(directory, "/", "_", -1) + ".wav",
			Mode: fuse.S_IFREG | 0644,
		}
	}

	return entries, fuse.OK
}

func wavDirAttr(_ *WavFsImpl, _ []string) (*fuse.Attr, fuse.Status) {
	return &fuse.Attr{
		Mode: fuse.S_IFDIR | 0755,
	}, fuse.OK
}
