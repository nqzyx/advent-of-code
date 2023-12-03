package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func printErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func getArgs(partName string) (dataPath string) {
	dataFolder := "../data"
	dataPath = path.Join(dataFolder, "input/data.txt")
	for _, arg := range os.Args[1:] {
		if arg == "-e" {
			dataPath = path.Join(dataFolder, "example", partName, "data.txt")
			break
		}
	}
	return
}

func getSchematic(partName string) (schematic []string) {
	dataPath := getArgs(partName)
	rawInput, err := os.ReadFile(dataPath)
	if err != nil {
		printErr(err)
	}
	schematic = strings.Split(string(rawInput), "\n")
	return
}

var specialCharacterRegexpString = "[^0-9.]"
var specialCharacterRegexp = regexp.MustCompile(specialCharacterRegexpString)

type specialCharacterPosition struct {
	row int
	col int
}

func findAllSpecialCharacters(schematic []string) (specialCharacterPositions []specialCharacterPosition) {
	for r, row := range schematic {
		for _, position := range specialCharacterRegexp.FindAllStringIndex(row, -1) {
			specialCharacterPositions = append(
				specialCharacterPositions,
				specialCharacterPosition{
					row: r,
					col: position[0],
				},
			)
		}
	}
	return
}

var serialNumberRegexpString = "[0-9]+"
var serialNumberRegexp = regexp.MustCompile(serialNumberRegexpString)

type serialNumberPosition struct {
	value    int
	row      int
	startCol int
	endCol   int // inclusive of last digit
}

func findAllSerialNumbers(schematic []string) (serialNumberPositions []serialNumberPosition) {
	for r, row := range schematic {
		for _, i := range serialNumberRegexp.FindAllStringIndex(row, -1) {
			serialNumber, err := strconv.Atoi(row[i[0]:i[1]])
			if err != nil {
				printErr(err)
			}
			serialNumberPositions = append(
				serialNumberPositions,
				serialNumberPosition{
					value:    serialNumber,
					row:      r,
					startCol: i[0],
					endCol:   i[1] - 1,
				},
			)
		}
	}
	return
}

func findTaggedSerialNumbers(schematic []string) (serialNumbers []int) {
	specialCharacterPositions := findAllSpecialCharacters(schematic)
	serialNumberPositions := findAllSerialNumbers(schematic)

	for _, serialNumberPosition := range serialNumberPositions {
		for _, specialCharacterPositions := range specialCharacterPositions {
			isTagged := specialCharacterPositions.row >= serialNumberPosition.row-1 &&
				specialCharacterPositions.row <= serialNumberPosition.row+1 &&
				specialCharacterPositions.col >= serialNumberPosition.startCol-1 &&
				specialCharacterPositions.col <= serialNumberPosition.endCol+1

			if isTagged {
				serialNumbers = append(serialNumbers, serialNumberPosition.value)
			}
		}
	}
	return
}

func partOne() (answer int) {
	schematic := getSchematic("part1")
	serialNumbers := findTaggedSerialNumbers(schematic)
	for _, serialNumber := range serialNumbers {
		answer += serialNumber
	}
	return
}

// func partTwo() (answer int) {
// 	input := getInput("part2")

// 	answer = len(input)
// 	return
// }

func main() {
	fmt.Printf("Part 1 Answer: %v\n", partOne())
	//fmt.Printf("Part 2 Answer: %v\n", partTwo())
}
