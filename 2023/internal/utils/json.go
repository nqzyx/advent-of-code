package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func JSONStringify(v any, indent bool) (string, error) {
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


func PrintlnJSON(v any, indent bool) (err error) {
	var jsonString string
	if jsonString, err = JSONStringify(v, indent); err != nil {
		return
	}
	fmt.Println(jsonString)
	return
}

func WriteJSONToFile(filename string, v any, indent bool) error {
	var jsonString string; var err error
	if jsonString, err = JSONStringify(v, indent); err != nil {
		return err
	}
	return WriteStringToFile(filename, jsonString)
}

type IOWriter io.Writer

func WriteJSON(w io.Writer, v any, indent bool) (err error) {
	var jsonString string
	if jsonString, err = JSONStringify(v, indent); err != nil {
		return
	}
	if _, err = w.Write([]byte(jsonString)); err != nil {
		return 
	}
	return
}