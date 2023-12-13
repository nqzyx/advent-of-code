package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/utils"
)

var digitsRegexp = regexp.MustCompile("[[:digit:]]+")

func getInput() (game Game) {
	ba, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	game = make(Game, 0, 5)
	dataStrings := strings.Split(string(ba), "\n")
	for _, line := range dataStrings {
		if len(line) == 0 {
			continue
		}
		tokens := strings.Split(line, " ")
		bid, _ := strconv.Atoi(tokens[1])
		game = append(game, NewHand(tokens[0], bid))
	}
	return
}

func partOne(game Game) int {
	game.RankHands()
	return game.CalculateWinnings()
}

func partTwo(game Game) int {
	return len(game)
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

func countWinners(r []int) int {
	var winners int
	raceTime, raceDistance := r[0], r[1]
	for holdTime := raceTime / 2; holdTime > 0; holdTime-- {
		speed := holdTime
		totalDistance := (raceTime - holdTime) * speed
		if totalDistance > raceDistance {
			winners++
		} else {
			break
		}
	}
	winners *= 2
	if raceTime%2 == 0 {
		winners--
	}
	return winners
}
