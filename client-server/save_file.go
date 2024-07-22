package client_server

import (
	"io"
	"os"
	"strings"
)

var tmpDir = "./tmp"

func saveFile(src io.Reader, filename string) error {
	if err := os.MkdirAll(tmpDir, 0750); err != nil {
		return err
	}

	file, err := os.Create(getFilename(filename))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, src)
	return err
}

func getFilename(filename string) string {
	return tmpDir + "/" + filename
}

func getFileExt(filename string) string {
	split := strings.Split(filename, ".")
	if len(split) == 0 {
		return ""
	}

	return split[len(split)-1]
}
