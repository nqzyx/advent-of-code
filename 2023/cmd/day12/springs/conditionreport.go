package springs

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type (
	ConditionReportRow struct {
		SpringConditions SpringConditions
		BrokenGroups     BrokenGroups
	}
	ConditionReport []ConditionReportRow
)

func NewConditionReport(crInput *[]string) *ConditionReport {
	report := make(ConditionReport, 0, len(*crInput))
	for _, crInRow := range *crInput {
		iParts := strings.Split(strings.TrimSpace(crInRow), " ")
		sc := *NewSpringConditions(iParts[0])
		bg := *NewBrokenGroups(iParts[1])
		report = append(report, ConditionReportRow{sc, bg})
	}
	return &report
}

func (crr *ConditionReportRow) MaxPadding() int {
	l := len(crr.SpringConditions)
	m := crr.BrokenGroups.MinLength()
	return l - m
}

func AddMatch(ma [][]int, m []int) [][]int {
	if !slices.ContainsFunc[[][]int, []int](ma, func(chk []int) bool {
		return slices.Equal(m, chk)
	}) {
		return append(ma, m)
	}
	return ma
}

func (crr *ConditionReportRow) MaxSolutions() (int, error) {
	re := regexp.MustCompile(crr.BrokenGroups.RegexpString())
	scs := string(crr.SpringConditions)
	fmt.Println("s:", scs, ", regex:", re.String())
	var matches [][]int
	for p := 0; p < crr.MaxPadding()+1; p++ {
		li, ri := p, len(scs)-p
		for _, s := range []string{scs[li:], scs[:ri]} {
			lma := re.FindAllStringIndex(s, -1)
			fmt.Printf("s: %v, ma: %v, matches: %v\n", s, lma, len(lma))
			for _, m := range lma {
				matches = AddMatch(matches, m)
			}
		}
		fmt.Println("For", scs, "Totals: matches:", len(matches))
	}
	return len(matches), nil
}
