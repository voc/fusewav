package fs

import (
	"regexp"

	"github.com/voc/fusewav/wav"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type path struct {
	regex  *regexp.Regexp
	attrFn func(fs *WavFsImpl, matches []string) (*fuse.Attr, fuse.Status)
	listFn func(fs *WavFsImpl, matches []string) ([]fuse.DirEntry, fuse.Status)
	readFn func(fs *WavFsImpl, matches []string) (file nodefs.File, code fuse.Status)
}

func compileOrNil(expr string) *regexp.Regexp {
	r, err := regexp.Compile(expr)
	if err != nil {
		return nil
	}

	return r
}

var paths = [...]path{
	{
		regex:  compileOrNil("^$"),
		attrFn: wavDirAttr,
		listFn: wavDirList,
		readFn: nil,
	},
	{
		regex:  compileOrNil("^([^/\\.]+\\.wav)$"),
		attrFn: wavFileAttr,
		listFn: nil,
		readFn: wavFileRead,
	},
}

// WavFsImpl represents the Implementation of all simulated directories and files inside the fuse-filesystem
type WavFsImpl struct {
	pathfs.FileSystem

	reader *wav.Reader
}

// GetAttr maps the attr syscall to the corresponding handler
func (fs *WavFsImpl) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	for _, path := range paths {
		matches := path.regex.FindStringSubmatch(name)
		if matches == nil {
			continue
		}

		if path.attrFn == nil {
			return nil, fuse.ENOENT
		}

		return path.attrFn(fs, matches)
	}

	return nil, fuse.ENOENT
}

// OpenDir maps the opendir syscall to the corresponding handler
func (fs *WavFsImpl) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	for _, path := range paths {
		matches := path.regex.FindStringSubmatch(name)
		if matches == nil {
			continue
		}

		if path.listFn == nil {
			return nil, fuse.ENOENT
		}

		return path.listFn(fs, matches)
	}

	return nil, fuse.ENOENT
}

// Open maps the open syscall to the corresponding handler
func (fs *WavFsImpl) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	for _, path := range paths {
		matches := path.regex.FindStringSubmatch(name)
		if matches == nil {
			continue
		}

		if flags&fuse.O_ANYWRITE != 0 {
			return nil, fuse.EPERM
		}

		if path.readFn == nil {
			return nil, fuse.ENOENT
		}

		return path.readFn(fs, matches)
	}

	return nil, fuse.ENOENT
}
