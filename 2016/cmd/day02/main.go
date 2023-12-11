package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func getInputData() []string {
	ba, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(ba), "\n")
	return s
}

func main() {
	directions := getInputData()
	fmt.Println(directions)

	answers, err := json.MarshalIndent(map[string]any{
		"Part 1": partOne(directions),
		"Part 2": partTwo(directions),
	}, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(answers))
}

func partOne(directions []string) int {
	return len(directions)
}

func partTwo(directions []string) int {
	return len(directions)
}
