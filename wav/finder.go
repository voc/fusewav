package wav

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

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

func findMatchingFiles(base string, directory string, start time.Time, end time.Time) ([]os.FileInfo, error) {
	matches := make([]os.FileInfo, 0)

	files, err := ioutil.ReadDir(filepath.Join(base, directory))
	if err != nil {
		return nil, err
	}

	var lastFileBefore os.FileInfo
	format := "2006-01-02_15-04-05"
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := file.Name()

		if filepath.Ext(filename) != ".wav" {
			continue
		}

		basename := filename[:len(filename)-4]
		time, err := time.Parse(format, basename)
		if err != nil {
			fmt.Printf("Wav-File with incompatible name skipped: %s", filename)
			continue
		}

		if time.Before(start) {
			lastFileBefore = file
			continue
		} else {
			if time.After(end) {
				if len(matches) > 0 {
					matches = append(matches, file)
				}
				break
			} else {
				if lastFileBefore != nil {
					matches = append(matches, lastFileBefore)
					lastFileBefore = nil
				}

				matches = append(matches, file)
			}
		}
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("No files in directory %s matched the requested timeframe %s - %s", directory, start, end)
	}

	return matches, nil
}
