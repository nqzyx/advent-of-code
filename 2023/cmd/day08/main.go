package main

import (
	"fmt"

	"nqzyx.xyz/advent-of-code/2023/day08/nodemap"
	"nqzyx.xyz/advent-of-code/2023/utils"
)

func partOne(input []string) (moves int) {
	nodeMap := nodemap.NewMap(input[1:])
	var err error
	if moves, err = nodeMap.CountMoves("AAA", "ZZZ", input[0]); err != nil {
		panic(err)
	}
	return moves
}

func partTwo(input []string) int {
	return len(input)
}

func main() {
	input := utils.GetInput()

	answers := map[string]any{
		"Part1": partOne(input),
		"Part2": partTwo(input),
	}

	utils.WriteJsonToFile("./answers.json", answers, true)
	fmt.Println(answers)
}
