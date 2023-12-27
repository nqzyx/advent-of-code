package main

import (
	"slices"
	"strings"

	"github.com/nqzyx/advent-of-code/2023/day11/galaxies"
	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
	pleaseIndentJSON := true
	input := utils.GetInput()

	// remove blank lines in input
	*input = slices.DeleteFunc(*input, func(inputRow string) bool {
		return len(strings.TrimSpace(inputRow)) == 0
	})

	u := galaxies.NewUniverse(input)

	answers := map[string]any{
		"part1": func() map[string]any {
			result, err := partOne(u)
			return map[string]any{
				"result": result,
				"err":    err,
			}
		}(),
		"part2": func() map[string]any {
			result, err := partTwo(u)
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

func partOne(u *galaxies.Universe) (int, error) {
	// fmt.Printf(banners["beginning"], "partOne")
	// defer fmt.Printf(banners["finished"], "partOne")

	td := int(u.GetTotalDistance())

	return td, nil
}

func partTwo(u *galaxies.Universe) (int, error) {
	// fmt.Printf(banners["beginning"], "partTwo")
	// defer fmt.Printf(banners["finished"], "partTwo")

	return len(*u), nil
}
