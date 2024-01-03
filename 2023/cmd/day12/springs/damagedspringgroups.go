package springs

import (
	"fmt"

	"github.com/nqzyx/advent-of-code/utils"
)

type (
	DamagedSpringGroupSize int
	DamagedSpringGroups    []DamagedSpringGroupSize
)

func NewDamagedSpringGroups(s string) *DamagedSpringGroups {
	dsgArr := utils.NewIntArray[int](s, ",")
	dsg := make(DamagedSpringGroups, len(dsgArr))
	for i, s := range dsgArr {
		dsg[i] = DamagedSpringGroupSize(s)
	}
	return &dsg
}

func (dsg DamagedSpringGroups) MinMatchLength() int {
	return int(utils.Sum(dsg...))
}

func (dsg DamagedSpringGroups) GroupMatchers() []string {
	matchers := make([]string, 0, len(dsg))
	for i, s := range dsg {
		matcher := fmt.Sprintf("(?P<dsg%v>%v{%v})", i+1, SC_DAMAGED.Matcher(), int(s))
		matchers = append(matchers, matcher)
	}
	return matchers
}

func (dsg DamagedSpringGroups) SeparatorMatcher() string {
	return fmt.Sprintf("(?P<sep%v>%v{%v})", "%v", SC_OPERATIONAL.Matcher(), "%v")
}
