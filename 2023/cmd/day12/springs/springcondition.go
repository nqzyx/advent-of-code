package springs

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type (
	SpringCondition     rune
	SpringConditionData struct {
		name    string
		matcher string
	}
)

const (
	SC_DAMAGED     SpringCondition = '#'
	SC_OPERATIONAL SpringCondition = '.'
	SC_UNKNOWN     SpringCondition = '?'
)

var springConditionDetails = map[SpringCondition]SpringConditionData{
	SC_OPERATIONAL: {
		name:    "SC_OPERATIONAL",
		matcher: fmt.Sprintf("[%v%v]", SC_OPERATIONAL, SC_UNKNOWN),
	},
	SC_DAMAGED: {
		name:    "SC_DAMAGED",
		matcher: fmt.Sprintf("[%v%v]", SC_DAMAGED, SC_UNKNOWN),
	},
	SC_UNKNOWN: {
		name:    "SC_UNKNOWN",
		matcher: fmt.Sprintf("[%v]", SC_UNKNOWN),
	},
}
var springConditionPrefix = "SC_"

func SpringConditions() (all []SpringCondition) {
	all = make([]SpringCondition, 0, len(springConditionDetails))
	all = append(all, maps.Keys(springConditionDetails)...)
	return
}

func NewSpringCondition(s string) SpringCondition {
	for _, sc := range SpringConditions() {
		if s == string(sc) || // one-character string constant
			s == sc.String() || // label w/o prefix
			s == sc.Name() { // label w/prefix
			return sc
		}
	}
	return SC_UNKNOWN
}

func NewSpringConditionArray(s string) *[]SpringCondition {
	sc := make([]SpringCondition, len(s))
	for i, s := range s {
		sc[i] = NewSpringCondition(string(s))
	}
	return &sc
}

func (sc SpringCondition) Matcher() string {
	return springConditionDetails[sc].matcher
}

// func (sc SpringCondition) MatcherRepeatInt(n int) string {
// 	repString := fmt.Sprintf("{%v}", n)
// 	return sc.MatcherRepeat(repString)
// }

// func (sc SpringCondition) MatcherRepeatRange(m, n int) string {
// 	repString := fmt.Sprintf("{%v:%v}", m, n)
// 	return sc.MatcherRepeat(repString)
// }

func (sc SpringCondition) Name() (scName string) {
	if scd, ok := springConditionDetails[sc]; ok {
		return scd.name
	} else {
		return string(sc)
	}
}

func (sc SpringCondition) String() string {
	if scd, ok := springConditionDetails[sc]; ok {
		return scd.name[len(springConditionPrefix):]
	}
	return string(sc)
}
