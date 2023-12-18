package utils

import (
	"encoding/json"
	"strings"
)

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
