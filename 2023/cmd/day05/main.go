package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/farmdata"
	"nqzyx.xyz/advent-of-code/2023/utils"
	"nqzyx.xyz/advent-of-code/2023/xref"
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
	farmData := farmdata.NewFarmData(getInputData(), true)
	if err := utils.WriteJsonToFile("./data/farmdata.json", farmData, true); err != nil {
		fmt.Println(err)
	}
	closestLocation := uint64(math.MaxUint64)
	for _, seed := range farmData.Seeds {
		if location, err := farmData.Resolve("seed", "location", seed); err == nil {
			closestLocation = min(location, closestLocation)
		} else {
			fmt.Println(err)
		}
	}
	answer = uint64(closestLocation)
	return
}

func processSeedRange(farmData *farmdata.FarmData, seedRange xref.Range, c chan<- uint64, index int) {
	minLocation := uint64(math.MaxUint64)
	fmt.Printf("%v: Seed range starting: %v\n", index, seedRange)
	for seed := seedRange.Start; seed < seedRange.End; seed++ {
		if seed%1000000 == 0 {
			fmt.Printf("%v: %v: %v (%v left)\n", index, seedRange, seed, (float32(seedRange.End-seed)/float32(seedRange.Length()))*100)
		}
		if location, err := farmData.Resolve("seed", "location", seed); err == nil {
			minLocation = min(location, minLocation)
		} else {
			fmt.Println(err)
		}
	}
	fmt.Printf("%v: Seed range finished: %v\n", index, seedRange)
	c <- minLocation
}

func partTwo() (answer uint64) {
	farmData := farmdata.NewFarmData(getInputData(), false)
	if err := utils.WriteJsonToFile("./data/farmdata.json", farmData, true); err != nil {
		fmt.Println(err)
	}
	closestLocation := uint64(math.MaxUint64)
	c := make(chan uint64, 1024)

	for r, rng := range farmData.SeedRanges {
		go processSeedRange(farmData, rng, c, r)
	}
	for range farmData.SeedRanges {
		closestLocation = min(<-c, closestLocation)
	}

	answer = uint64(closestLocation)
	return
}

func main() {
	answersJson, err := utils.JsonStringify(map[string]any{
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
