package springs

import (
	"fmt"
	"strings"

	"github.com/nqzyx/advent-of-code/utils"
)

type (
	BrokenGroupSize int
	BrokenGroups    []BrokenGroupSize
)

func NewBrokenGroups(s string) *BrokenGroups {
	dgsArr := utils.NewNumericArrayFromStringWithSeparator[int](s, ",")
	dgs := make(BrokenGroups, len(dgsArr))
	for i, s := range dgsArr {
		dgs[i] = BrokenGroupSize(s)
	}
	return &dgs
}

func (bg *BrokenGroups) MinLength() int {
	// Start with number of separators
	ml := len(*bg) - 1
	// Add the size of each group
	for _, s := range *bg {
		ml += int(s)
	}
	return ml
}

func (bg *BrokenGroups) RegexpStrings(sc SpringCondition) []string {
	sa := make([]string, 0, len(*bg))
	if len(*bg) == 0 {
		return sa
	}
	for _, s := range *bg {
		sa = append(sa, sc.RegexpStringRepeatInt(int(s)))
	}
	return sa
}

func (bg *BrokenGroups) RegexpString() string {
	// ALL separators are based on SC_OPERATIONAL
	// (i.e., finding a good well)
	// The outer separator (ends) should appear 0-1
	// times. We prefer 0, if possible
	oSep := SC_OPERATIONAL.RegexpString() + "??"
	// The inner separator (middles) should appear 1
	// or more times. We prefer 1, if possible
	iSep := SC_OPERATIONAL.RegexpString() + "+?"
	// Now, construct the regexp string for this row...
	return fmt.Sprintf("%v%v%v",
		// start with a separator
		oSep,
		// join the group regexp strings with sep
		strings.Join(bg.RegexpStrings(SC_DAMAGED), iSep),
		// end with a separator
		oSep,
	)
}

func (bg *BrokenGroups) SeparatorCounts() []int {
	c := len(*bg) + 1
	ia := make([]int, c)
	if c == 0 {
		return ia
	}
	for i := 0; i < c; i++ {
		if i == 0 || i == c-1 {
			continue
		}
		ia[i] = 1
	}
	return ia
}
