package wav

import (
	"fmt"
	"os"
	"time"
)

// The Reader reads and assembles segments of the named wav-files within the given time frame
type Reader struct {
	base                string
	directories         []string
	directoryToFilesMap map[string][]os.FileInfo
}

// NewReader constructs a new Wav-Reader instance
func NewReader(base string, start time.Time, end time.Time) (*Reader, error) {
	fmt.Printf("Initializing reader for segments in %s between %s and %s\n", base, start, end)

	directoriesToConsider, err := findDirectories(base)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Considering directories: %s\n", directoriesToConsider)

	directories := make([]string, 0)
	directoryToFilesMap := make(map[string][]os.FileInfo)
	for _, directory := range directoriesToConsider {
		files, err := findMatchingFiles(base, directory, start, end)
		if err != nil {
			return nil, err
		}

		if len(files) == 0 {
			continue
		}

		filenames := make([]string, len(files))
		for i, file := range files {
			filenames[i] = file.Name()
		}
		fmt.Printf("Matching files in directory %s: %s\n", directory, filenames)
		directoryToFilesMap[directory] = files
		directories = append(directories, directory)
	}

	reader := &Reader{
		base:                base,
		directories:         directories,
		directoryToFilesMap: directoryToFilesMap,
	}
	return reader, nil
}

// ListDirectories matches lists all Directories in the base-dir which have audio-files in the requested range
func (reader *Reader) ListDirectories() []string {
	return reader.directories
}

// GetAggregatedWavFile creates a new AggregatedWavFile representing all Wav-Files in the Time-Window in the given directory
func (reader *Reader) GetAggregatedWavFile(directory string) (*AggregatedWavFile, error) {
	files, ok := reader.directoryToFilesMap[directory]
	if !ok {
		return nil, fmt.Errorf("Unknown directory %s", directory)
	}

	return NewAggregatedWavFile(reader.base, directory, files), nil
}
