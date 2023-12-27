package main

import (
	"fmt"

	"github.com/nqzyx/advent-of-code/2023/day09/oasis"
	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
	input := utils.GetInput()
	report := oasis.NewReport(*input)

	answers := make(map[string]any, 2)

	if result, err := partOne(report); err == nil {
		answers["part1"] = result
	} else {
		answers["part1"] = err
	}
	if result, err := partTwo(report); err == nil {
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

func partOne(report *oasis.Report) (moves int, err error) {
	// fmt.Println("partOne: Beginning")
	// defer fmt.Println("partOne: Finished")

	// utils.WriteJSONToFile("./report.json", report, true)
	return report.PredictionTotals.Next, nil
}

func partTwo(report *oasis.Report) (moves int, err error) {
	// fmt.Println("partTwo: Beginning")
	// defer fmt.Println("partTwo: Finished")

	// utils.WriteJSONToFile("./report.json", report, true)
	return report.PredictionTotals.Previous, nil
}
