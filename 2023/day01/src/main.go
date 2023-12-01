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

var digitNamesToValues map[string]int64 = map[string]int64 { 
	"one": 1, 
	"two": 2, 
	"three": 3, 
	"four": 4, 
	"five": 5, 
	"six": 6, 
	"seven": 7, 
	"eight": 8, 
	"nine": 9,
}
var digitNameReplacements map[string]string = map[string]string { 
	"one": "o1e", 
	"two": "t2o", 
	"three": "t3e", 
	"four": "f4r",
	"five": "f5e", 
	"six": "s6x",
	"seven": "s7n", 
	"eight": "e8t", 
	"nine": "n9e",
}

var notDigitsRegexpString string = "[^1-9]"
var notDigitsRegexp regexp.Regexp = *regexp.MustCompile(notDigitsRegexpString)

func mapInputToNumsByDigits(strs []string) (nums []int64) {
	for _, str := range strs {
		str = notDigitsRegexp.ReplaceAllString(str, "")
		num, _ := strconv.ParseInt(str[0:1] + str[len(str) -1:], 10, 64)
		nums = append(nums, num)
	}
	return
}

func findNum(str string) (num int64, found bool) {
	ch := str[:1]
	if ch >= "0" && ch <= "9" {
		num, _ = strconv.ParseInt(ch, 10, 64)
		found = true
		return
	}

	for name, value := range digitNamesToValues {
		if strings.HasPrefix(str, name) {
			num = value
			found = true
			return
		}
	}
	return 0, false
}

func mapInputToNumsBySubstr(strs []string) (nums []int64) {
	for _, str := range strs {
		var a, b int64
		var found bool
		// Find first num
		for idx := range str {
			if a, found = findNum(str[idx:]); found {
				break
			}
		}

		// Find last num by searching in reverse
		for idx := len(str) - 1; idx >= 0; idx-- {
			if b, found = findNum(str[idx:]); found {
				break
			}
		}
		nums = append(nums, (a * 10) + b)
	}
	return
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
	nums := mapInputToNumsBySubstr(input)
	for _, n := range nums {
		total += int64(n)
	}
	return
}

func main() {
	fmt.Printf("Part 1 Answer: %v\n", partOne())
	fmt.Printf("Part 2 Answer: %v\n", partTwo())
}
