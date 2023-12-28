package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nqzyx/advent-of-code/2023/day12/springs"
	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
	pleaseIndentJSON := true
	input := utils.GetInput()

	// remove blank lines in input
	*input = slices.DeleteFunc(*input, func(inputRow string) bool {
		return len(strings.TrimSpace(inputRow)) == 0
	})

	answers := map[string]any{
		"part1": func() map[string]any {
			result, err := partOne(input)
			return map[string]any{
				"result": result,
				"err":    err,
			}
		}(),
		"part2": func() map[string]any {
			result, err := partTwo(input)
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
	"finished":  ">>>------> %v: Finished <------<<<\n\n",
}

func partOne(input *[]string) (int, error) {
	fmt.Printf(banners["beginning"], "partOne")
	defer fmt.Printf(banners["finished"], "partOne")

	crs := springs.NewReport(input)

	fmt.Println("crs:", crs)

	return len(*crs), nil
}

func partTwo(input *[]string) (int, error) {
	fmt.Printf(banners["beginning"], "partTwo")
	defer fmt.Printf(banners["finished"], "partTwo")

	crs := springs.NewReport(input)

	fmt.Println("crs:", crs)

	return len(*crs), nil
}
