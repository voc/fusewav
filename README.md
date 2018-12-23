# FuseWav
A Fuse-Filesystem for Zero-Copy concatenation of PCM-Wav-Files

## Install
For macOS install [FUSE for macOS](https://osxfuse.github.io/) and run `/Library/Filesystems/osxfuse.fs/Contents/Resources/load_osxfuse`.

For Ubuntu/Debian install `libfuse2`.

Install GoLang: https://golang.org/dl/

Run `go get github.com/voc/fusewav`. The compiles binary now is in `$GOPATH/bin/fusewav`.

## Usage
```
# For macOS
/Library/Filesystems/osxfuse.fs/Contents/Resources/load_osxfuse

fusewav --start "2018-12-22 23:00" --end "2018-12-23 00:30" --base /video/audio-backups --mountpoint ~/Desktop/Rescue "pa/l" "pa/r" "headset/*" "hand/*" &

# Work with ~/Desktop/Rescue
/v/audio-backup ls -la ~/Desktop/Rescue/
total 269568
drwxr-xr-x  1 pkoerner  staff         0  1 Jan  1970 .
drwx------+ 8 pkoerner  staff       256 23 Dez 02:05 ..
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 hand_0.wav
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 hand_1.wav
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 hand_2.wav
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 headsets_0.wav
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 headsets_1.wav
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 headsets_2.wav
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 pa_l.wav
-rw-r--r--  1 pkoerner  staff  17250236  1 Jan  1970 pa_r.wav

umount ~/Desktop/Rescue
```

# Expected Folder-Layout
Something like this (compatible with https://github.com/voc/aes67-recorder)
```
./hand/0/2018-12-22_20-00-36.wav
./hand/0/2018-12-22_19-59-36.wav
./hand/0/2018-12-22_19-59-06.wav
./hand/0/2018-12-22_20-00-06.wav
./hand/1/2018-12-22_20-00-36.wav
./hand/1/2018-12-22_19-59-36.wav
./hand/1/2018-12-22_19-59-06.wav
./hand/1/2018-12-22_20-00-06.wav
./hand/2/2018-12-22_20-00-36.wav
./hand/2/2018-12-22_19-59-36.wav
./hand/2/2018-12-22_19-59-06.wav
./hand/2/2018-12-22_20-00-06.wav
./headsets/0/2018-12-22_20-00-36.wav
./headsets/0/2018-12-22_19-59-36.wav
./headsets/0/2018-12-22_19-59-06.wav
./headsets/0/2018-12-22_20-00-06.wav
./headsets/1/2018-12-22_20-00-36.wav
./headsets/1/2018-12-22_19-59-36.wav
./headsets/1/2018-12-22_19-59-06.wav
./headsets/1/2018-12-22_20-00-06.wav
./headsets/2/2018-12-22_20-00-36.wav
./headsets/2/2018-12-22_19-59-36.wav
./headsets/2/2018-12-22_19-59-06.wav
./headsets/2/2018-12-22_20-00-06.wav
./pa/r/2018-12-22_20-00-36.wav
./pa/r/2018-12-22_19-59-36.wav
./pa/r/2018-12-22_19-59-06.wav
./pa/r/2018-12-22_20-00-06.wav
./pa/l/2018-12-22_20-00-36.wav
./pa/l/2018-12-22_19-59-36.wav
./pa/l/2018-12-22_19-59-06.wav
./pa/l/2018-12-22_20-00-06.wav
```