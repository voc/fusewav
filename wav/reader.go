package wav

import (
	"fmt"
	"path/filepath"
	"time"
)

// The Reader reads and assembles segments of the named wav-files within the given time frame
type Reader struct {
	base        string
	start       time.Time
	end         time.Time
	directories []string
}

// NewReader constructs a new Wav-Reader instance
func NewReader(base string, start time.Time, end time.Time, patterns []string) (*Reader, error) {
	fmt.Printf("Initializing reader for segments in %s matching %s between %s and %s\n", base, patterns, start, end)

	directories, err := findDirectoriesMatching(base, patterns)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Matching directories: %s\n", directories)

	reader := &Reader{
		base:        base,
		start:       start,
		end:         end,
		directories: directories,
	}
	return reader, nil
}

// ListMatchingDirectories matches the given patterns against the existing Directories and returns matching ones
func (reader *Reader) ListMatchingDirectories() []string {
	return reader.directories
}

func findDirectoriesMatching(base string, patterns []string) ([]string, error) {
	matches := make([]string, 0)
	for _, pattern := range patterns {
		pathForPattern := filepath.Join(base, pattern)
		matchesForPattern, _ := filepath.Glob(pathForPattern)

		if len(matchesForPattern) == 0 {
			return nil, fmt.Errorf("No directories for pattern %s found in basedir %s", pattern, base)
		}

		for _, match := range matchesForPattern {
			relativeMatch, _ := filepath.Rel(base, match)
			matches = append(matches, relativeMatch)
		}
	}

	return matches, nil
}
