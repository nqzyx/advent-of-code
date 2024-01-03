package springs

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nqzyx/advent-of-code/utils"
)

type (
	ConditionReportRow struct {
		SpringConditions    []SpringCondition
		DamagedSpringGroups DamagedSpringGroups
	}
	ConditionReport []ConditionReportRow
)

func NewConditionReport(crInput *[]string) *ConditionReport {
	report := make(ConditionReport, 0, len(*crInput))
	for _, crInputRow := range *crInput {
		crRow := NewConditionReportRow(crInputRow)
		report = append(report, *crRow)
	}
	return &report
}

func NewConditionReportRow(crrInput string) *ConditionReportRow {
	iParts := strings.Split(strings.TrimSpace(crrInput), " ")
	return &ConditionReportRow{
		*NewSpringConditionArray(iParts[0]),
		*NewDamagedSpringGroups(iParts[1]),
	}
}

func (crr *ConditionReportRow) MaxPadding() int {
	return utils.Max(len(string(crr.SpringConditions))-crr.DamagedSpringGroups.MinMatchLength(), 0)
}

func (crr *ConditionReportRow) Matches() []string {
	conditionReports := string(crr.SpringConditions)
	groupMatchers := crr.DamagedSpringGroups.GroupMatchers()
	separatorMatcher := crr.DamagedSpringGroups.SeparatorMatcher()
	maxPadding := crr.MaxPadding()
	separatorCount := len(groupMatchers) + 1

	done := false

	matcherSB := new(strings.Builder)
	for g, groupMatcher := range groupMatchers {
		if g == 0 {
			matcherSB.WriteString(fmt.Sprintf(separatorMatcher, g, "%v"))
		}
		matcherSB.WriteString(groupMatcher)
		matcherSB.WriteString(fmt.Sprintf(separatorMatcher, g+1, "%v"))
	}
	rowMatcher := matcherSB.String()

	genSetSeed := make([]uint, 0, separatorCount)
	for sc := 0; sc<separatorCount; sc++ {
		seed := uint(1)
		if sc == 0 || sc == separatorCount -1 {
			seed = uint(0)
		}
		genSetSeed = append(genSetSeed, seed)
	} // result will be [0, 1, ... 1, 0]
	genSetSeedArgs := []any{genSetSeed}

	var matches = make([]string, 0, utils.SumOfIntegers(maxPadding)*separatorCount)
	var matchFmt = "%v [%v:%v]"

	fmt.Println("maxPadding:", maxPadding)

	if maxPadding == 0 { // no need to iterate...
		matcher := fmt.Sprintf(rowMatcher, genSetSeedArgs...)
		loc := regexp.MustCompile(matcher).FindStringIndex(conditionReports)
		if len(loc) != 0 {
			for m := 0; m < len(loc); m += 2 {
				matches = append(matches, fmt.Sprintf(matchFmt, conditionReports[loc[m]:loc[m+1]], loc[m], loc[m+1]))
			}
		}
	} else {
		generatorSet := utils.MustNewGeneratorSetSeeded[uint](
			genSetSeed,
			uint(maxPadding),
			func(gs *utils.GeneratorCollection[uint]) { done = true },
		)
		for separatorRepeatCounts := generatorSet.Next(); !done;  {
			// fmt.Println("repeatCounts:", separatorRepeatCounts)
			separatorRepeatCountsAny := []any{separatorRepeatCounts}
			matcher := fmt.Sprintf(rowMatcher, separatorRepeatCountsAny...) 
			loc := regexp.MustCompile(matcher).FindStringIndex(conditionReports)
			if len(loc) != 0 {
				for m := 0; m < len(loc); m += 2 {
					matches = append(matches, fmt.Sprintf("%v [%v:%v]", conditionReports[loc[m]:loc[m+1]], loc[m], loc[m+1]))				
				}
			}
		}
	}
	return matches
}
