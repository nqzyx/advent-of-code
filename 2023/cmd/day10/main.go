package main

import (
	"fmt"

	"github.com/nqzyx/advent-of-code/day10/pipes"
	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
	input := utils.GetInput()

	var pipeMap *pipes.Map
	var err error

	if pipeMap, err = pipes.NewMap(&input); err != nil {
		panic(err)
	}

	fmt.Println(pipeMap)

	answers := map[string]any{
		"part1": func() map[string]any {
			result, err := partOne(pipeMap)
			return map[string]any{
				"result": result,
				"err":    err,
			}
		}(),
		"part2": func() map[string]any {
			result, err := partTwo(pipeMap)
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

func partOne(m *pipes.Map) (length int, err error) {
	// fmt.Println("partOne: Beginning")
	// defer fmt.Println("partOne: Finished")

	// utils.PrintlnJSON(m, true)

	utils.WriteJSONToFile("./map.json", *m, true)

	return m.PathLength(), nil
}

func partTwo(m *pipes.Map) (length int, err error) {
	// fmt.Println("partTwo: Beginning")
	// defer fmt.Println("partTwo: Finished")

	length = len(*m.Insiders)
	fmt.Println("Path:", *m.Path)
	fmt.Println("Insiders:", *m.Insiders)

	// utils.PrintlnJSON(*m.Insiders, false)
	return
}
