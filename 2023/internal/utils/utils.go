package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

var integerRegexp = regexp.MustCompile("[[:digit:]]+")

func Parse[T constraints.Integer](s string, v *T) (err error) {
	_, err = fmt.Sscanf(s, "%v", v)
	return
}

func NewIntArrayFromString[T constraints.Integer](s string) (result []T) {
	s = strings.TrimSpace(s)
	sArr := integerRegexp.FindAllString(s, -1)
	result = make([]T, len(sArr))
	for i, v := range sArr {
		var x T
		err := Parse(v, &x)
		if err != nil {
			fmt.Println(err)
		}
		result[i] = T(x)
	}
	return
}

func JsonStringify(v any, indent bool) (string, error) {
	const (
		ASCII_SPACE rune = 0x20
	)
	var jsonBytes []byte
	var jsonString string
	var err error
	if indent {
		jsonBytes, err = json.MarshalIndent(v, "", strings.Repeat(string(ASCII_SPACE), 2))
	} else {
		jsonBytes, err = json.Marshal(v)
	}
	if err != nil {
		return "", err
	}
	jsonString = string(jsonBytes)
	return jsonString, nil
}

func WriteJsonToFile(filename string, v any, indent bool) error {
	var jsonString string
	var err error
	if jsonString, err = JsonStringify(v, indent); err != nil {
		return err
	}
	return WriteStringToFile(filename, jsonString)
}

func WriteStringToFile(filename string, content string) error {
	var file *os.File
	var err error

	openFlags := os.O_CREATE + os.O_WRONLY

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
