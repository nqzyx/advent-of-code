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
