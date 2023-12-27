package main

import (
	"fmt"

	"github.com/nqzyx/advent-of-code/2023/day10/pipes"
	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
	pleaseIndentJSON := true
	inputData := utils.GetInput()

	var m *pipes.Map
	var err error

	if m, err = pipes.NewMap(inputData); err != nil {
		panic(err)
	}

	utils.PrintlnJSON(m, pleaseIndentJSON)

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


	utils.WriteJSONToFile("./answers.json", answers, pleaseIndentJSON)
	utils.PrintlnJSON(answers, pleaseIndentJSON)
}

var banners = map[string]string{
	"beginning": ">>>------> %v: Beginning <------<<<\n\n",
	"finished": ">>>------> %v: Finished <------<<<\n\n",
}

func partOne(m *pipes.Map) (int, error) {
	fmt.Printf(banners["beginning"], "partOne")
	defer fmt.Printf(banners["finished"], "partOne")

	fmt.Printf("MAP\n---\n%v\n\n", m)

	return m.PathLength(), nil
}

func partTwo(m *pipes.Map) (int, error) {
	fmt.Printf(banners["beginning"], "partTwo")
	defer fmt.Printf(banners["finished"], "partTwo")

	length := len(m.Insiders) 

	return length, nil
}
