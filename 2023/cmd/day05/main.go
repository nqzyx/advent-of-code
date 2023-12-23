package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"strings"

	"github.com/nqzyx/advent-of-code/day05/almanac"
	"github.com/nqzyx/advent-of-code/utils"
)

func getArgs() (dataPath string) {
	dataBaseFolder := "./data"
	dataPath = path.Join(dataBaseFolder, "input", "data.txt")
	args := os.Args[1:]
	for a, arg := range args {
		if arg == "-e" {
			dataPath = path.Join(dataBaseFolder, "example", fmt.Sprintf("data%v.txt", args[a+1]))
			break
		}
	}
	return
}

func getInputData() (inputData []string) {
	dataPath := getArgs()
	inputDataAsByteArray, err := os.ReadFile(dataPath)
	if err != nil {
		panic(err)
	}
	inputData = strings.Split(string(inputDataAsByteArray), "\n\n")
	return
}

func partOne() (answer uint64) {
	almanac := almanac.NewAlmanac(getInputData(), true)
	if err := utils.WriteJSONToFile("./data/almanac.json", almanac, true); err != nil {
		fmt.Println(err)
	}
	closestLocation := uint64(math.MaxUint64)
	for _, seed := range almanac.Seeds {
		if location, err := almanac.Resolve("seed", "location", seed); err == nil {
			closestLocation = min(location, closestLocation)
		} else {
			fmt.Println(err)
		}
	}
	answer = uint64(closestLocation)
	return
}

func partTwo() (answer uint64) {
	almanac := almanac.NewAlmanac(getInputData(), false)
	if err := utils.WriteJSONToFile("./data/almanac.json", almanac, true); err != nil {
		fmt.Println(err)
	}
	closestLocation := uint64(math.MaxUint64)
	for _, rng := range almanac.SeedRanges {
		for seed := rng.Start; seed < rng.End; seed++ {
			if location, err := almanac.Resolve("seed", "location", seed); err == nil {
				closestLocation = min(location, closestLocation)
			} else {
				fmt.Println(err)
			}
		}
	}
	answer = uint64(closestLocation)
	return
}

func main() {
	answersJson, err := utils.JSONStringify(map[string]any{
		"Part 1": partOne(),
		"Part 2": partTwo(),
	}, true)
	if err == nil {
		err = utils.WriteStringToFile("./data/answers.json", answersJson)
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answersJson)
}
