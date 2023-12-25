package oasis

import (
	"github.com/nqzyx/advent-of-code/utils"
)

type ValueHistory struct {
	Readings    []int
	Analyses    Analyses
	Predictions struct {
		Previous int
		Next     int
	}
}

func NewValueHistory(valueList string) (result *ValueHistory) {
	readings := utils.NewNumericArrayFromString[int](valueList)
	result = &ValueHistory{
		Readings: readings,
	}
	result.Analyze()
	return
}

func (h *ValueHistory) setPredictions() {
	h.Predictions.Next = h.Analyses[0].Predictions.Next
	h.Predictions.Previous = h.Analyses[0].Predictions.Previous
}

func (vh *ValueHistory) Analyze() {
	if vh == nil || len(vh.Readings) == 0 {
		return
	}

	vh.Analyses = make(Analyses, 0, len(vh.Readings))
	currentAnalysis := NewAnalysisInitialized(vh.Readings)
	vh.Analyses = append(vh.Analyses, currentAnalysis)

	for {
		if nextAnalysis := currentAnalysis.NextAnalysis(); nextAnalysis == nil {
			break
		} else {
			vh.Analyses = append(vh.Analyses, nextAnalysis)
			currentAnalysis = NewAnalysisInitialized(nextAnalysis.Intervals)
		}
	}

	vh.Analyses.Analyze()

	vh.setPredictions()

	// utils.PrintlnJSON(vh, false)
}
