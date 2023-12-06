package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func printErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func getArgs() (dataPath string) {
	dataBaseFolder := "../data"
	dataPath = path.Join(dataBaseFolder, "input", "data.txt")
	args := os.Args[1:]
	for a, arg := range args {
		if arg == "-e" {
			dataPath = path.Join(dataBaseFolder, "example", fmt.Sprintf("data%v.txt", args[a+1]))
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
	inputData = strings.Split(string(inputDataAsByteArray), "\n\n")
	return
}

type GameData struct {
	seeds   []int
	xRefMap map[string]XRef
}

var digitsRegexp = regexp.MustCompile("([1-9][0-9]+)")

func getSeedList(str string) []int {
	strings.Trim()
}

func parseInputData(inputData []string) GameData {
	gameData := new(GameData)
	for _, data := range inputData {
		switch true {
		case strings.HasPrefix(data, "seeds:"):
			gameData.seeds = getSeedList(data)
		default:
			gameData.addXrefMap(data)
		}
	}
	return gameData
}

func partOne() (answer int) {
	inputData := getInputData()

	fmt.Println(inputData)

	xref := NewXRef("seed", "soil").AddXRefRange(5, 43, 49)
	fmt.Println(xref)

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
