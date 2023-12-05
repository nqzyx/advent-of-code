package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func printErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func getArgs() (dataPath string) {
	dataBaseFolder := "../data"
	dataPath = path.Join(dataBaseFolder, "input", "data.txt")
	for _, arg := range os.Args[1:] {
		if arg == "-e" {
			dataPath = path.Join(dataBaseFolder, "example", "data.txt")
			break
		}
	}
	return
}

func getInputData() (inputData []string) {
	dataPath := getArgs()
	inputDataAsByteArray, err := os.ReadFile(dataPath)
	if err != nil {
		printErr(err)
	}
	inputData = strings.Split(string(inputDataAsByteArray), "\n")
	return
}

func partOne() (answer int) {
	inputData := getInputData()

	// Do the needful

	answer = len(inputData)
	return
}

func partTwo() (answer int) {
	inputData := getInputData()

	// Do the needful

	answer = len(inputData)
	return
}

func main() {
	fmt.Printf("Part 1 Answer: %v\n", partOne())
	fmt.Printf("Part 2 Answer: %v\n", partTwo())
}
