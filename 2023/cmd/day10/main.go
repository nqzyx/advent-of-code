package main

import (
	"fmt"

	"github.com/nqzyx/advent-of-code/2023/day10/pipes"
	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
	input := utils.GetInput()

	var m *pipes.Map
	var err error

	if m, err = pipes.NewMap(&input); err != nil {
		panic(err)
	}

	// fmt.Println("Map:", m)

	answers := map[string]any{
		"part1": func() map[string]any {
			result, err := partOne(m)
			return map[string]any{
				"result": result,
				"err":    err,
			}
		}(),
		"part2": func() map[string]any {
			result, err := partTwo(m)
			return map[string]any{
				"result": result,
				"err":    err,
			}
		}(),
	}

	pleaseIndent := true

	utils.WriteJSONToFile("./answers.json", answers, pleaseIndent)
	utils.PrintlnJSON(answers, pleaseIndent)
}

var beginning string = ">>>------> %v: Beginning <------<<<\n\n"
var finished string = ">>>------> %v: Finished <------<<<\n\n"

func partOne(m *pipes.Map) (int, error) {
	fmt.Printf(beginning, "partOne")
	defer fmt.Printf(finished, "partOne")

	fmt. Println(m)

	return m.PathLength(), nil
}

func partTwo(m *pipes.Map) (int, error) {
	fmt.Printf(beginning, "partTwo")
	defer fmt.Printf(finished, "partTwo")

	length := len(*m.Insiders)

	fmt.Printf("Path: %v\n\n", *m.Path)

	fmt.Printf("Insiders: %v\n\n", *m.Insiders)

	return length, nil
}
