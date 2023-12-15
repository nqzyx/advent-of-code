package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/day07/camelcards"
	"nqzyx.xyz/advent-of-code/2023/utils"
)

func getInput() []string {
	ba, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(ba), "\n")
}

func partOne(input []string) int64 {
	jokerRule := false
	game := camelcards.NewGame(input, jokerRule)
	utils.WriteJsonToFile("part1_game.json", game, true)
	return game.CalculateWinnings()
}

func partTwo(input []string) int64 {
	jokerRule := true
	game := camelcards.NewGame(input, jokerRule)
	utils.WriteJsonToFile("part2_game.json", game, true)
	return game.CalculateWinnings()
}

func main() {
	input := getInput()

	answersJson, err := utils.JsonStringify(map[string]any{
		"Part 1": partOne(input),
		"Part 2": partTwo(input),
	}, true)
	if err == nil {
		err = utils.WriteStringToFile("./answers.json", answersJson)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answersJson)
}
