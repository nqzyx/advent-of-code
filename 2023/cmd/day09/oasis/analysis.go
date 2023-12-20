package oasis

type Analysis struct{
	Intervals []int
	Predictions struct{
		Previous int
		Next int
	}
}

type Analyses []*Analysis

func NewAnalysis(length int) (a *Analysis) {
	return &Analysis{Intervals: make([]int, length)}
}

func NewAnalysisInitialized(intervals []int) (a *Analysis) {
	a = NewAnalysis(len(intervals))
	copy(a.Intervals, intervals)
	return 
}

func (prev Analysis) NextAnalysis() (next *Analysis) {
	if prev.isFinal() {
		return nil
	}
	next = &Analysis{ Intervals: make([]int, 0, len(prev.Intervals) - 1) }
	for i := 0; i < len(prev.Intervals) -1; i++ {
		next.Intervals = append(next.Intervals, prev.Intervals[i+1]-prev.Intervals[i])
	}
	return
}

func (a Analysis) isFinal() (result bool) {
	if a.Intervals == nil || len(a.Intervals) == 0 {
		return true
	}
	for _, value := range a.Intervals[1:] {
		if value != a.Intervals[0] {
			return false
		}
	}
	return true
}

func (aList *Analyses) Analyze() (prev int, next int) {
	if aList == nil || len(*aList) == 0 {
		return
	}
	lastIndex := len(*aList) - 1
	for i := lastIndex; i >= 0; i-- {
		thisA := (*aList)[i]
		if i == lastIndex {
			thisA.Predictions.Previous = thisA.Intervals[0]
			thisA.Predictions.Next = thisA.Intervals[len(thisA.Intervals) -1]
		} else {
			nextA := (*aList)[i + 1]
			thisA.Predictions.Previous = thisA.Intervals[0] - nextA.Predictions.Previous
			thisA.Predictions.Next = thisA.Intervals[len(thisA.Intervals)-1] + nextA.Predictions.Next
		}
		prev += thisA.Predictions.Previous
		next += thisA.Predictions.Next
	}
	return
}