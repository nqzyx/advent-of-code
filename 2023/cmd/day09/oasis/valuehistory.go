package oasis

import "nqzyx.xyz/advent-of-code/2023/utils"

type ValueHistory struct {
	Readings   []int
	Analysis   [][]int
	Prediction int
}

func NewValueHistory(valueList string) *ValueHistory {
	readings := utils.NewIntArrayFromString[int](valueList)
	return &ValueHistory{
		Readings: readings,
		Analysis: make([][]int, len(readings)/2),
	}
}
