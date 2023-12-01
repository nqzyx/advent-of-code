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
			break;
		}
	}
	return
}

func getInput(partName string) (input []string) {
	dataPath := getArgs(partName)
	rawInput, err := os.ReadFile(dataPath)
	if err != nil {
		printErr(err)
	}
	input = strings.Split(string(rawInput), "\n")
	return
}

var digitNamesToValues map[string]int64 = map[string]int64 { "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9 }
var digitNameReplacements map[string]string = map[string]string {  "one": "o1e", "two": "t2o", "three": "th3ee", "five": "fi5ve", "seven": "se7en", "eight": "ei8ht", "nine": "ni9ne"}

var digitNamesRegexpString string = "(one|two|three|four|five|six|seven|eight|nine)"
var digitNamesRegexp regexp.Regexp = *regexp.MustCompile(digitNamesRegexpString)

var digitsRegexpString string = "([1-9])"
var digitsRegexp regexp.Regexp = *regexp.MustCompile(digitsRegexpString)

var notDigitsRegexpString string = "[^1-9]"
var notDigitsRegexp regexp.Regexp = *regexp.MustCompile(notDigitsRegexpString)

func replaceDigitName (str string) (string) {
	return digitNameReplacements[str]
}		

func convertNamesToDigits (str string) (result string) {
	lastString := str
	for {
		thisString := digitNamesRegexp.ReplaceAllStringFunc(lastString,replaceDigitName)
		if thisString == lastString {
			result = lastString
			return
		}
		lastString = thisString
	}
}

func mapInputToNumsByDigits(strs []string) (nums []int64) {
	for _, str := range strs {
		str = notDigitsRegexp.ReplaceAllString(str, "")
		num, _ := strconv.ParseInt(str[0:1] + str[len(str) -1:], 10, 64)
		nums = append(nums, num)
	}
	return
}

func mapInputToNumsByDigitsAndNames(strs []string) ([]int64) { 
	var tmpStrs []string
	for _, str := range strs {
		tmpStrs = append(tmpStrs, convertNamesToDigits(str))
	}
	return mapInputToNumsByDigits(tmpStrs)
}

func partOne() (total int64) {
	input := getInput("part1")

	nums := mapInputToNumsByDigits(input)
	for _, n := range nums {
		total += int64(n)
	}
	return
}



func partTwo() (total int64) {
	input := getInput("part2")
	nums := mapInputToNumsByDigitsAndNames(input)

	for _, n := range nums {
		total += int64(n)
	}
	return
}

func main() {
	// s := "oneighthreeightwonesevenine"
	// fmt.Printf("{\n\tstring:\t%v,\n\tresult:\t%v\n}", s, convertNamesToDigits(s))
	fmt.Printf("Part 1 Answer: %v\n", partOne())
	fmt.Printf("Part 2 Answer: %v\n", partTwo())
}
