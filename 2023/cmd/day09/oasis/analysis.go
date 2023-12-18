package oasis

type Analysis []int

type Analyses []Analysis

func (a Analysis) isFinalAnalysis() (result bool) {
	if len(a) == 0 {
		return true
	}
	initialValue := a[0]
	result = true
	for _, i := range a[1:] {
		result = result && i == initialValue
	}
	return
}

func (a Analysis) getNextAnalysis() (next Analysis, isFinal bool) {
	if len(a) == 0 || a.isFinalAnalysis() {
		return nil, true
	}

	next = make(Analysis, 0, len(a)-1)

	for i := 0; i < len(a)-1; i++ {
		next = append(next, a[i+1]-a[i])
	}
	return
}
