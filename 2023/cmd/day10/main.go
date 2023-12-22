package main

import (
	"fmt"

	"nqzyx.xyz/advent-of-code/2023/day10/pipes"
	"nqzyx.xyz/advent-of-code/2023/utils"
)

func main() {
	input := utils.GetInput()

	var pipeMap *pipes.Map
	var err error

	if pipeMap, err = pipes.NewMap(input); err != nil {
		panic(err)
	}

	answers := map[string]any{
		"part1": func() any {
			if result, err := partOne(pipeMap); err != nil {
				return map[string]any{
					"err": err,
				}
			} else {
				return map[string]any{
					"result": result,
				}
			}
		}(),
		"part2": func() any {
			if result, err := partTwo(pipeMap); err != nil {
				return err
			} else {
				return result
			}
		}(),
	}

	pleaseIndent := true

	utils.WriteJSONToFile("./answers.json", answers, pleaseIndent)
	utils.PrintlnJSON(answers, pleaseIndent)
}

func partOne(m *pipes.Map) (length int, err error) {
	fmt.Println("partOne: Beginning")
	defer fmt.Println("partOne: Finished")

	// utils.PrintlnJSON(m, true)

	// utils.WriteJSONToFile("./report.json", report, true)
	if length, err := m.FindPathLength(); err != nil {
		return 0, err
	} else {
		return length, nil
	}
}

func partTwo(m *pipes.Map) (length int, err error) {
	// fmt.Println("partTwo: Beginning")
	// defer fmt.Println("partTwo: Finished")

	// utils.WriteJSONToFile("./report.json", report, true)
	return m.Rows * m.Cols, nil
}
