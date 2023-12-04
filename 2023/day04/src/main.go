package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func printErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func getArgs() (dataPath string) {
	dataBaseFolder := "../data"
	args := os.Args[1:]
	dataPath = path.Join(dataBaseFolder, "input", "data.txt")
	for i, arg := range args {
		if arg == "-e" {
			dataPath = path.Join(dataBaseFolder, "example", fmt.Sprintf("data%v.txt", args[i+1]))
			break
		}
	}
	return
}

type InputData []string

func getInputData() (inputData InputData) {
	dataPath := getArgs()
	inputDataAsByteArray, err := os.ReadFile(dataPath)
	if err != nil {
		printErr(err)
	}
	inputData = strings.Split(strings.Trim(strings.ReplaceAll(string(inputDataAsByteArray), "Card ", ""), " "), "\n")
	return
}

type String string
type StringArray []string
type IntArray []int
type Card struct {
	score          int
	winningNumbers IntArray
	theNumbers     IntArray
}
type Cards map[int]Card

func (s String) ToIntArray(splitOn string) (result IntArray) {
	this := strings.TrimSpace(string(s))
	space, space2 := " ", "  "
	for strings.Contains(this, space2) {
		this = strings.ReplaceAll(this, space2, space)
	}
	theseStrings := strings.Split(this, splitOn)
	result = make(IntArray, len(theseStrings))
	for thisIndex, thisString := range theseStrings {
		thisValue, err := strconv.Atoi(thisString)
		if err != nil {
			fmt.Printf("thisIndex: %v, thisString: %v, this; %v", thisIndex, thisString, this)
			printErr(err)
		}
		result[thisIndex] = thisValue
	}
	return
}

func (ia IntArray) Contains(value int) (result bool) {
	result = false
	for _, i := range ia {
		if i == value {
			result = true
			break
		}
	}
	return
}

func (c Card) CalculateScore() (result int) {
	for _, n := range c.theNumbers {
		if c.winningNumbers.Contains(n) {
			if result == 0 {
				result = 1
			} else {
				result *= 2
			}
		}
	}
	return
}

func (input InputData) ToCards() (cards Cards) {
	cards = make(Cards)
	var thisCardData []string
	for thisCardIndex, thisCardString := range input {
		thisCardData = strings.Split(strings.TrimSpace(thisCardString), ":")
		thisCardNumber, err := strconv.Atoi(thisCardData[0])
		if err != nil {
			fmt.Printf("thisCardIndex: %v, thisCardData[0]: %v, thisCardString: %v", thisCardIndex, thisCardData[0], thisCardString)
			printErr(err)
		}
		thisCardData = strings.Split(thisCardData[1], "|")
		thisCard := Card{
			score:          0,
			winningNumbers: String(thisCardData[0]).ToIntArray(" "),
			theNumbers:     String(thisCardData[1]).ToIntArray(" "),
		}
		thisCard.score = thisCard.CalculateScore()
		cards[thisCardNumber] = thisCard
	}
	return
}

func partOne() (answer int) {
	cards := getInputData().ToCards()
	for _, c := range cards {
		answer += c.score
	}
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
