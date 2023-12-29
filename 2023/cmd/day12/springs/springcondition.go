package springs

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type (
	SpringCondition        rune
	SpringConditions       []SpringCondition
	springConditionDetails struct {
		name         string
		regexpString string
	}
)

const (
	SC_DAMAGED     SpringCondition = '#'
	SC_OPERATIONAL SpringCondition = '.'
	SC_UNKNOWN     SpringCondition = '?'
)

var springConditions = map[SpringCondition]springConditionDetails{
	SC_OPERATIONAL: {"SC_OPERATIONAL", fmt.Sprintf("[%v%v]", string(SC_OPERATIONAL), string(SC_UNKNOWN))},
	SC_DAMAGED:     {"SC_DAMAGED", fmt.Sprintf("[%v%v]", string(SC_DAMAGED), string(SC_UNKNOWN))},
	SC_UNKNOWN:     {"SC_UNKNOWN", fmt.Sprintf("[%v]", string(SC_UNKNOWN))},
}
var springConditionPrefix = "SC_"

func allSpringConditions() (all SpringConditions) {
	all = make(SpringConditions, 0, len(springConditions))
	all = append(all, maps.Keys(springConditions)...)
	return
}

func NewSpringCondition(s string) SpringCondition {
	for _, sc := range allSpringConditions() {
		if s == string(sc) || // one-character string constant
			s == sc.String() || // label w/o prefix
			s == sc.Name() { // label w/prefix
			return sc
		}
	}
	return SC_UNKNOWN
}

func NewSpringConditions(s string) *SpringConditions {
	sc := make(SpringConditions, len(s))
	for i, s := range s {
		sc[i] = NewSpringCondition(string(s))
	}
	return &sc
}

func (sc SpringCondition) RegexpString() (regexp string) {
	return springConditions[sc].regexpString
}

func (sc SpringCondition) RegexpStringRepeat(r string) string {
	return fmt.Sprintf("(%v)%v", springConditions[sc].regexpString, r)
}

func (sc SpringCondition) RegexpStringRepeatInt(n int) string {
	repString := fmt.Sprintf("{%v}", n)
	return sc.RegexpStringRepeat(repString)
}

func (sc SpringCondition) RegexpStringRepeatRange(m, n int) string {
	repString := fmt.Sprintf("{%v:%v}", m, n)
	return sc.RegexpStringRepeat(repString)
}

func (sc SpringCondition) Name() (scName string) {
	if scd, ok := springConditions[sc]; ok {
		return scd.name
	} else {
		return string(sc)
	}
}

func (sc SpringCondition) String() string {
	if scd, ok := springConditions[sc]; ok {
		return scd.name[len(springConditionPrefix):]
	} else {
		return string(sc)
	}
}
