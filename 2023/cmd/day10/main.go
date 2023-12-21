package main

import (
	"fmt"

	"nqzyx.xyz/advent-of-code/2023/day10/pipes"
	"nqzyx.xyz/advent-of-code/2023/utils"
)

func main() {
	input := utils.GetInput()
	pipeMap := pipes.NewMap(input)

	answers := make(map[string]any, 2)

	if result, err := partOne(pipeMap); err == nil {
		answers["part1"] = result
	} else {
		answers["part1"] = err
	}
	if result, err := partTwo(input); err == nil {
		answers["part2"] = result
	} else {
		answers["part2"] = err
	}

	pleaseIndent := true

	if answersAsJson, err := utils.JSONStringify(answers, pleaseIndent); err != nil {
		panic(err)
	} else {
		utils.WriteStringToFile("./answers.json", answersAsJson)
		fmt.Println(answersAsJson)
	}
}

func partOne(m *pipes.Map) (moves int, err error) {
	fmt.Println("partOne: Beginning")
	utils.PrintlnJSON(m, true)
	defer fmt.Println("partOne: Finished")

	// utils.WriteJSONToFile("./report.json", report, true)
	return m.GetPathLength(), nil
}

func partTwo(input []string) (moves int, err error) {
	// fmt.Println("partTwo: Beginning")
	// defer fmt.Println("partTwo: Finished")

	// utils.WriteJSONToFile("./report.json", report, true)
	return len(input), nil
}
