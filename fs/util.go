package fs

import (
	"strings"

	"github.com/voc/fusewav/wav"
)

func filenameFromDirectoryName(directoryName string) string {
	return strings.Replace(directoryName, "/", "_", -1) + ".wav"
}

func findDirectoryForFilename(reader *wav.Reader, filename string) string {
	for _, directory := range reader.ListDirectories() {
		if filename == filenameFromDirectoryName(directory) {
			return directory
		}
	}

	return ""
}
