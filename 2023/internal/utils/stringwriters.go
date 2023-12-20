package utils

import (
	"os"
)

func WriteStringToFile(filename string, content string) error {
	var file *os.File
	var err error

	openFlags := os.O_CREATE + os.O_TRUNC + os.O_WRONLY

	if file, err = os.OpenFile(filename, openFlags, os.FileMode(0o777)); err != nil {
		return err
	}
	if _, err = file.WriteString(content); err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	return nil
}
