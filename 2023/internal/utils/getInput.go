package utils

import (
	"io"
	"os"
	"strings"
)

func GetInput() []string {
	ba, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(ba), "\n")
}
