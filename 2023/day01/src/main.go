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

func mapInputToNums(strs []string) (nums []int) {
	for _, str := range strs {
		notDigits, err := regexp.Compile("[^0-9]")
		if err != nil {
			printErr(err)
		}
		str = notDigits.ReplaceAllString(str, "")
		num, _ := strconv.Atoi(string(str[0]) + string(str[len(str) - 1]))
		nums = append(nums, num)
	}
	return
}

func partOne() {
	input := getInput("part1")

	nums := mapInputToNums(input)

	var total int
	for _, n := range nums {
		total += n
	}
	fmt.Println("total", total)
}



func partTwo() {

}

func main() {
	partOne()
	partTwo()
}