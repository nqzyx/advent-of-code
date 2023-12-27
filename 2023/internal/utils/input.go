package utils

import (
	"io"
	"os"
	"strings"
)

func GetInput() *[]string {
	ba, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(ba), "\n")
	return &s
}
