package file

import (
	"io"
	"os"
)

func CopyFile(toFile, fromFile string) (n int64, err error) {
	fromfile, err := os.Open(fromFile)
	if err != nil {
		return
	}
	defer fromfile.Close()

	tofile, err := os.Create(toFile)
	if err != nil {
		return
	}
	defer tofile.Close()

	return io.Copy(tofile, fromfile)
}
