package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/nqzyx/advent-of-cde2023/utils"
)

var digitsRegexp = regexp.MustCompile("[[:digit:]]+")

func getInput() []int {
	ba, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	dataStrings := strings.Split(string(ba), "\n")
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(dataStrings[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(dataStrings[1], ":")[1], " ", ""))
	var game []int = []int{
		time,
		distance,
	}
	return game
}

func partOne(races []int) int {
	var answer int = 1
	// for _, r := range races {
	// 	answer *= countWinners()
	// 	// fmt.Println("Game:", i, ", Winners:", winners)
	// }
	return answer
}

func partTwo(race []int) int {
	return countWinners(race)
	// return len(race)
}

func main() {
	input := getInput()
	answersJson, err := utils.JSONStringify(map[string]any{
		"Part 1": partOne(input),
		"Part 2": partTwo(input),
	}, true)
	if err == nil {
		err = utils.WriteStringToFile("./data/answers.json", answersJson)
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
