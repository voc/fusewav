package wav

import (
	"fmt"
	"os"
	"time"
)

const headerSize = 40

// The Reader reads and assembles segments of the named wav-files within the given time frame
type Reader struct {
	directories         []string
	directoryToFilesMap map[string][]os.FileInfo
}

// NewReader constructs a new Wav-Reader instance
func NewReader(base string, start time.Time, end time.Time, patterns []string) (*Reader, error) {
	fmt.Printf("Initializing reader for segments in %s matching %s between %s and %s\n", base, patterns, start, end)

	directories, err := findDirectoriesMatching(base, patterns)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Matching directories: %s\n", directories)

	directoryToFilesMap := make(map[string][]os.FileInfo)
	for _, directory := range directories {
		files, err := findMatchingFiles(base, directory, start, end)
		if err != nil {
			return nil, err
		}

		filenames := make([]string, len(files))
		for i, file := range files {
			filenames[i] = file.Name()
		}
		fmt.Printf("Matching files in directory %s: %s\n", directory, filenames)
		directoryToFilesMap[directory] = files
	}

	reader := &Reader{
		directories:         directories,
		directoryToFilesMap: directoryToFilesMap,
	}
	return reader, nil
}

// ListMatchingDirectories matches the given patterns against the existing Directories and returns matching ones
func (reader *Reader) ListMatchingDirectories() []string {
	return reader.directories
}

func (reader *Reader) GetAggregatedWavFile(directory string) (*AggregatedWavFile, error) {
	files, ok := reader.directoryToFilesMap[directory]
	if !ok {
		return nil, fmt.Errorf("Unknown directory %s", directory)
	}

	return NewAggregatedWavFile(files), nil
}
