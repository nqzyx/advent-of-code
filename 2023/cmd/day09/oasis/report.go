package oasis

import (
	"strings"
)

type Report struct {
	ValueHistories  []*ValueHistory
	PredictionTotals struct {
		Next int
		Previous int
	}
}

func NewReport(reportInput []string) (newReport *Report) {
	valueHistories := make([]*ValueHistory, 0, len(reportInput))
	for _, valueHistory := range reportInput {
		if len(strings.TrimSpace(valueHistory)) == 0 {
			continue
		}
		valueHistory := NewValueHistory(valueHistory)
		valueHistories = append(valueHistories, valueHistory)
	}
	newReport = &Report{
		ValueHistories: valueHistories,
	}

	for _, h := range newReport.ValueHistories {
		newReport.PredictionTotals.Next += h.Predictions.Next
		newReport.PredictionTotals.Previous += h.Predictions.Previous
	}

	return
}
