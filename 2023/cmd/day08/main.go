package main

import (
	"fmt"

	"nqzyx.xyz/advent-of-code/2023/day08/nodemap"
	"nqzyx.xyz/advent-of-code/2023/utils"
)

func partOne(input []string)  (moves int, err error) {
	if debug {
		fmt.Println("partOne: Beginning")  
		defer fmt.Println("partOne: Finished")
	}

	nodeMap := nodemap.New(input[1:])
	var startingNodeNames []string
	startingNodeNames, err = nodeMap.FindMatchingNodeNames("A$")
	if err != nil {
		return 0, err
	}
	moves, err = nodeMap.PathLength(input[0], startingNodeNames[0], "Z$")
	if err != nil {
		return 0, err
	}
	return moves, nil
}

func partTwo(input []string) (moves int, err error) {
	if debug {
		fmt.Println("partTwo: Beginning")
		defer fmt.Println("partTwo: Finished")
	}

	nodeMap := nodemap.New(input[1:])
	if moves, err = nodeMap.FindMultiplePathLengths(input[0], "A$", "Z$"); err != nil {
		return 0, err
	}
	return moves, nil
}
var debug bool

func main() {
	debug = false
	input := utils.GetInput()
	answers := make(map[string]map[string]any, 2);

	var answer map[string]any

	answer = make(map[string]any)
	if moves, err := partOne(input); err == nil {
		answer["moves"] = moves
	} else {
		answer["err"] = err
	}
	answers["part1"] = answer

	answer = make(map[string]any)
	if moves, err := partTwo(input); err == nil {
		answer["moves"] = moves
	} else {
		answer["err"] = err
	}
	answers["part2"] = answer

	pleaseIndent := true
	if answersAsJson, err := utils.JsonStringify(answers, pleaseIndent); err != nil {
		panic(err)
	} else {
		utils.WriteStringToFile("./answers.txt", answersAsJson)
		fmt.Println(answersAsJson)
	}
}
