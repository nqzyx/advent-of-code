package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

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
		panic(err)
	}
	inputData = strings.Split(strings.TrimSpace(strings.ReplaceAll(string(inputDataAsByteArray), "Card ", "")), "\n")
	return
}

type String string
type StringArray []string
type IntArray []int
type Card struct {
	winningNumbers   IntArray
	candidateNumbers IntArray
	winners          int // how many candidates match winners
	count            int // how many instances of this card
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
			panic(err)
		}
		result[thisIndex] = thisValue
	}
	return
}

func (ia IntArray) contains(value int) (result bool) {
	result = false
	for _, i := range ia {
		if i == value {
			result = true
			break
		}
	}
	return
}

func (c Card) getWinners() (result int) {
	for _, candidate := range c.candidateNumbers {
		if c.winningNumbers.contains(candidate) {
			result += 1
		}
	}
	return
}

func (c Card) String() (result string) {
	format := "{\n\twinningNumbers: %v,\n\tcandidateNumbers: %v,\n\twinners: %v\n\tcount: %v\n}"
	result = fmt.Sprintf(
		format,
		c.winningNumbers,
		c.candidateNumbers,
		c.winners,
		c.count,
	)
	return
}

func (c Card) score() (result int) {
	if c.winners > 0 {
		result = int(math.Pow(float64(2), float64(c.winners-1)))
	}
	return
}
func (input InputData) toCards() (cards Cards) {
	cards = make(Cards)
	var thisCardData []string
	for _, thisCardString := range input {
		thisCardData = strings.Split(strings.TrimSpace(thisCardString), ":")
		thisCardNumber, err := strconv.Atoi(thisCardData[0])
		if err != nil {
			panic(err)
		}
		thisCardData = strings.Split(thisCardData[1], "|")
		thisCard := Card{
			winningNumbers:   String(thisCardData[0]).ToIntArray(" "),
			candidateNumbers: String(thisCardData[1]).ToIntArray(" "),
			count:            1,
		}
		thisCard.winners = thisCard.getWinners()
		cards[thisCardNumber] = thisCard
	}
	for c := 1; c <= len(cards); c++ {
		thisCard := cards[c]
		for i := 1; i <= thisCard.winners; i++ {
			if nextCard, ok := cards[c+i]; ok {
				nextCard.count += thisCard.count
				cards[c+i] = nextCard
			}
		}
	}
	return
}

func partOne() (answer int) {
	cards := getInputData().toCards()
	for _, c := range cards {
		answer += c.score()
	}
	return
}

func partTwo() (answer int) {
	cards := getInputData().toCards()
	for _, c := range cards {
		answer += c.count
	}
	return
}

func main() {
	answers := map[string]int{
		"Part1": partOne(),
		"Part2": partTwo(),
	}

	jsonBytes, err := json.MarshalIndent(answers, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}
