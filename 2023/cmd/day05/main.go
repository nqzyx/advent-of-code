package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/farmdata"
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
	farmData := farmdata.NewFarmData(getInputData())
	fmt.Println(json.Marshal(farmData))
	closestLocation, _ := strconv.ParseUint("0xFFFFFFFFFFFFFFFF", 0, 64)
	for _, seed := range farmData.Seeds() {
		var location uint64
		location, err := farmData.DestinationValueByType("seed", seed, "location")
		if err != nil {
			panic(err)
		}
		closestLocation = min(location, closestLocation)
	}
	answer = uint64(closestLocation)
	return
}

func partTwo() (answer uint64) {
	inputData := getInputData()

	// Do the needful

	answer = uint64(len(inputData))
	return
}

func main() {
	fmt.Println("Part 1:", partOne())
	fmt.Println("Part 2:", partTwo())
}
