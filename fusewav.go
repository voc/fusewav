package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/voc/fusewav/fs"
	"github.com/voc/fusewav/wav"
)

const exitCodeConfigError = 1
const exitCodeRuntimeError = 2

func main() {
	base := flag.String("base", "/video/audio-backup", "The root-folder of the recorded segments.")
	start := flag.String("start", "", "The first used segment will be the one which contains this Point in Time. Format: 2018-12-23_00-46-35")
	end := flag.String("end", "", "The last used segment will be the one which contains this Point in Time. Format: 2018-12-23_00-46-35")
	mountpoint := flag.String("mountpoint", "", "Path to mount the assembled view of all segments.")
	flag.Parse()

	if *base == "" || *start == "" || *end == "" || *mountpoint == "" {
		fmt.Printf("Required Fields: base, start, end\n")
		flag.Usage()
		os.Exit(exitCodeConfigError)
	}

	format := "2006-01-02 15:04"
	startDate, errStart := time.Parse(format, *start)
	endDate, errEnd := time.Parse(format, *end)
	if errStart != nil || errEnd != nil {
		fmt.Printf("Start- or End-Date format was incorrect, required: 2018-12-23 00:46\n")
		flag.Usage()
		os.Exit(exitCodeConfigError)
	}

	reader, err := wav.NewReader(*base, startDate, endDate, flag.Args())
	if err != nil {
		fmt.Printf("Error constructing Wav-Reader: %s", err)
		os.Exit(exitCodeConfigError)
	}

	fs, err := fs.NewWavFs(reader, *mountpoint)
	if err != nil {
		fmt.Printf("Error constructing Wav-Fs: %s", err)
		os.Exit(exitCodeRuntimeError)
	}

	fs.Serve()
}
