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
var gearRegexpString = "\\*"

type specialCharacterPosition struct {
	row int
	col int
}

func findAllSpecialCharacters(schematic []string, searchFor string) (specialCharacterPositions []specialCharacterPosition) {
	specialCharacterRegexp := regexp.MustCompile(searchFor)
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
	specialCharacterPositions := findAllSpecialCharacters(schematic, specialCharacterRegexpString)
	serialNumberPositions := findAllSerialNumbers(schematic)

	for _, serialNumberPosition := range serialNumberPositions {
		for _, specialCharacterPosition := range specialCharacterPositions {
			isTagged := specialCharacterPosition.row >= serialNumberPosition.row-1 &&
				specialCharacterPosition.row <= serialNumberPosition.row+1 &&
				specialCharacterPosition.col >= serialNumberPosition.startCol-1 &&
				specialCharacterPosition.col <= serialNumberPosition.endCol+1

			if isTagged {
				serialNumbers = append(serialNumbers, serialNumberPosition.value)
			}
		}
	}
	return
}

func findGearRatios(schematic []string) (gearRatios []int) {
	for _, gearPosition := range findAllSpecialCharacters(schematic, gearRegexpString) {
		var gearSerialNumbers []int
		for _, serialNumber := range findAllSerialNumbers(schematic) {
			isAdjacentToGear := gearPosition.row >= serialNumber.row-1 &&
				gearPosition.row <= serialNumber.row+1 &&
				gearPosition.col >= serialNumber.startCol-1 &&
				gearPosition.col <= serialNumber.endCol+1
			if isAdjacentToGear {
				gearSerialNumbers = append(gearSerialNumbers, serialNumber.value)
			}
		}
		if len(gearSerialNumbers) == 2 {
			gearRatios = append(gearRatios, gearSerialNumbers[0]*gearSerialNumbers[1])
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

func partTwo() (answer int) {
	schematic := getSchematic("part2")
	gearRatios := findGearRatios(schematic)

	for _, gearRatio := range gearRatios {
		answer += gearRatio
	}
	return
}

func main() {
	fmt.Printf("Part 1 Answer: %v\n", partOne())
	fmt.Printf("Part 2 Answer: %v\n", partTwo())
}
