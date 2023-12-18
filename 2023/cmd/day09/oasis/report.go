package oasis

type Report struct {
	ValueHistories  []*ValueHistory
	PredictionTotal int64
}

func NewReport(report []string) *Report {
	valueHistories := make([]*ValueHistory, len(report))
	for i, valueHistoryList := range report {
		valueHistory := NewValueHistory(valueHistoryList)
		valueHistories[i] = valueHistory
	}
	return &Report{
		ValueHistories: valueHistories,
	}
}
