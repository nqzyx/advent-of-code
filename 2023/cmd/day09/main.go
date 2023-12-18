package main

import (
	"fmt"

	"nqzyx.xyz/advent-of-code/2023/utils"
)

func main() {
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

func partOne(input []string)  (moves int, err error) {
	// fmt.Println("partOne: Beginning")  
	// defer fmt.Println("partOne: Finished")

	return len(input), nil
}

func partTwo(input []string) (moves int, err error) {
	// fmt.Println("partTwo: Beginning")
	// defer fmt.Println("partTwo: Finished")

	return len(input), nil
}
