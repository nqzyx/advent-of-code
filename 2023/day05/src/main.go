package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func jsonPrint(v any) {
	indent := strings.Repeat(" ", 2)
	ba, err := json.MarshalIndent(v, "", indent)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ba))
}

func getArgs() (dataPath string) {
	dataBaseFolder := "../data"
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

func partOne() (answer uint32) {
	gardenData := NewGardenData(getInputData())
	jsonPrint(map[string]interface{}{"gardenData": gardenData})
	closest, _ := strconv.ParseUint("0xFFFFFFFF", 0, 32)
	for _, seed := range gardenData.seeds {
		location := uint64(gardenData.GetTargetValue("seed", seed, "location"))
		if location < closest {
			closest = location
		}
	}
	answer = uint32(closest)
	return
}

func partTwo() (answer uint32) {
	inputData := getInputData()

	// Do the needful

	answer = uint32(len(inputData))
	return
}

func main() {
	jsonPrint(map[string]uint32{
		"Part1": partOne(),
		"Part2": partTwo(),
	})
}
