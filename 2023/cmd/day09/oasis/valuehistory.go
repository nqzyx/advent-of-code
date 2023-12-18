package oasis

import (
	"fmt"

	"nqzyx.xyz/advent-of-code/2023/utils"
)

type ValueHistory struct {
	Readings   []int
	Analyses   Analyses
	Prediction int
}

func NewValueHistory(valueList string) (result *ValueHistory) {
	readings := utils.NewIntArrayFromString[int](valueList)
	result = &ValueHistory{
		Readings: readings,
	}
	result.Analyze()
	return
}

func (h *ValueHistory) setPrediction() {
	fmt.Println("valueHistory:", h.Analyses)
	h.Prediction = 0
	for _, a := range h.Analyses {
		h.Prediction += a[len(a)-1]
	}
}

func (h *ValueHistory) Analyze() {
	if h == nil || len(h.Readings) == 0 {
		return
	}
	h.Analyses = make(Analyses, 0, len(h.Readings)/2)

	h.Analyses = append(h.Analyses, Analysis(h.Readings))
	current := h.Analyses[0]

	for {
		var final bool
		if current, final = current.getNextAnalysis(); !final {
			h.Analyses = append(h.Analyses, current)
		} else {
			break
		}
	}

	h.setPrediction()
}
