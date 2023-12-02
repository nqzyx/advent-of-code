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

type Draw map[string]int
type Game []Draw
type Board map[int]Game

func getBoard(partName string) (board Board) {
	board = make(Board)
	dataPath := getArgs(partName)
	rawInput, err := os.ReadFile(dataPath)
	if err != nil {
		printErr(err)
	}
	boardData := strings.Split(string(rawInput), "\n")
	for _, gameData := range boardData {
		gameParts := strings.Split(gameData, ":")
		gameNumber, err := strconv.Atoi(strings.Split(gameParts[0], " ")[1])
		if err != nil {
			panic(fmt.Sprintf("Game number not found (%v)", gameData))
		}
		drawsData := strings.Split(gameParts[1], ";")
		thisGame := make(Game, 0, 0)
		for _, drawData := range drawsData {
			thisDraw := make(Draw)
			cubesData := strings.Split(drawData, ",")
			for _, cubeData := range cubesData {
				cubeCountData := strings.Split(strings.Trim(cubeData, " "), " ")
				// fmt.Println(cubeCountData)
				cubeCountQty, _ := strconv.Atoi(cubeCountData[0])
				cubeCountColor := cubeCountData[1]
				thisDraw[cubeCountColor] = cubeCountQty
			}
			thisGame = append(thisGame, thisDraw)
		}
		board[gameNumber] = thisGame
	}

	return
}



// var digitNamesToValues map[string]int64 = map[string]int64 { 
// 	"one": 1, 
// 	"two": 2, 
// 	"three": 3, 
// 	"four": 4, 
// 	"five": 5, 
// 	"six": 6, 
// 	"seven": 7, 
// 	"eight": 8, 
// 	"nine": 9,
// }
// var digitNameReplacements map[string]string = map[string]string { 
// 	"one": "o1e", 
// 	"two": "t2o", 
// 	"three": "t3e", 
// 	"four": "f4r",
// 	"five": "f5e", 
// 	"six": "s6x",
// 	"seven": "s7n", 
// 	"eight": "e8t", 
// 	"nine": "n9e",
// }

// var notDigitsRegexpString string = "[^1-9]"
// var notDigitsRegexp regexp.Regexp = *regexp.MustCompile(notDigitsRegexpString)

// func mapInputToNumsByDigits(strs []string) (nums []int64) {
// 	for _, str := range strs {
// 		str = notDigitsRegexp.ReplaceAllString(str, "")
// 		num, _ := strconv.ParseInt(str[0:1] + str[len(str) -1:], 10, 64)
// 		nums = append(nums, num)
// 	}
// 	return
// }

// func findNum(str string) (num int64, found bool) {
// 	ch := str[:1]
// 	if ch >= "0" && ch <= "9" {
// 		num, _ = strconv.ParseInt(ch, 10, 64)
// 		found = true
// 		return
// 	}

// 	for name, value := range digitNamesToValues {
// 		if strings.HasPrefix(str, name) {
// 			num = value
// 			found = true
// 			return
// 		}
// 	}
// 	return 0, false
// }

// func mapInputToNumsBySubstr(strs []string) (nums []int64) {
// 	for _, str := range strs {
// 		var a, b int64
// 		var found bool
// 		// Find first num
// 		for idx := range str {
// 			if a, found = findNum(str[idx:]); found {
// 				break
// 			}
// 		}

// 		// Find last num by searching in reverse
// 		for idx := len(str) - 1; idx >= 0; idx-- {
// 			if b, found = findNum(str[idx:]); found {
// 				break
// 			}
// 		}
// 		nums = append(nums, (a * 10) + b)
// 	}
// 	return
// }

var colors = []string { "red", "green", "blue" }

func partOne() (total int) {
	board := getBoard("part1")
	limits := map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	for g, game := range board {
		possible := true
		top:
		for _, draw := range game {
			for _, color := range colors {
				if draw[color] > limits[color] {
					possible = false
					break top
				}
			}
		}
		if possible {
			total += g
		}
	}
	return
}



func partTwo() (total int) {
	board := getBoard("part2")
	for _, game := range board {
		minCubesRequired := make(map[string]int)
		for _, draw := range game {
			for _, color := range colors {
				if minCubesRequired[color] < draw[color] {
					minCubesRequired[color] = draw[color]
				}
 			}
		}
		power := 1
		for _, minQty := range minCubesRequired {
			power *= minQty
		}
		total += power
	}
	return
}

func main() {
	fmt.Printf("Part 1 Answer: %v\n", partOne())
	fmt.Printf("Part 2 Answer: %v\n", partTwo())
}
