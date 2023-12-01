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

func mapToNums(strs []string) (nums []int) {
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

func getArgs() (dataPath string) {
	dataFolder := "../data"
	dataPath = path.Join(dataFolder, "input/data.txt")
	for _, arg := range os.Args[1:] {
		if arg == "-e" {
			dataPath = path.Join(dataFolder, "example/data.txt")
			break;
		}
	}
	return
}

func partOne() {
	dataPath := getArgs()
	rawInput, err := os.ReadFile(dataPath)
	if err != nil {
		printErr(err)
	}
	inputArray := strings.Split(string(rawInput), "\n")
	nums := mapToNums(inputArray)

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