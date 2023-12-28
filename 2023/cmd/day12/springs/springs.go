package springs

import (
	"fmt"
	"strings"

	"github.com/nqzyx/advent-of-code/utils"
)

type (
	ConditionReport struct {
		Springs          string
		DamageGroupSizes []int
	}
	ConditionReports []ConditionReport
)

func NewReport(input *[]string) (conditionReports *ConditionReports) {
	conditionReports = new(ConditionReports)
	*conditionReports = make(ConditionReports, 0, len(*input))
	for _, iRow := range *input {
		inRowParts := strings.Split(strings.TrimSpace(iRow), " ")
		springs, damagedSpringGroupSizes := inRowParts[0], utils.NewNumericArrayFromStringWithSeparator[int](inRowParts[1], ",")
		*conditionReports = append(*conditionReports, ConditionReport{springs, damagedSpringGroupSizes})
	}
	return
}

func GetMaximumMatchLength(groupSizeList []int) int {
	maxMatchLength := len(groupSizeList) - 1
	for _, size := range groupSizeList {
		maxMatchLength += size
	}
	return maxMatchLength
}

func GetGroupMatchString(groupMatchString string, lastGroup bool, separatorMatchString string) (string, error) {
	if lastGroup {
		return groupMatchString, nil
	} else {
		return groupMatchString + separatorMatchString, nil
	}
}

func GetMatchStringListForGroups(groups []int, separatorMatchString string) (matchStrings []string, err error) {
	matchStrings = make([]string, 0)
	if len(groups) == 0 {
		return
	}
	groupsMatchStrings := make([]string, 0, len(groups))
	var groupMatchString string
	for gIndex, gSize := range groups {
		groupMatchString, err = GetGroupMatchString(SC_DAMAGED.GetSpringGroupMatchString(gSize), gIndex == len(groups)-1, separatorMatchString)
		if err != nil {
			return nil, err
		}
		groupsMatchStrings = append(groupsMatchStrings, groupMatchString)
	}
}

func GetReportMatchStrings(springs string, groupSizeList []int) (reportMatchStringList []string, err error) {
	reportMatchStringList = make([]string, 0)
	separatorMatchString := SC_OPERATIONAL.GetSpringGroupMatchString(1)
	maximumMatchLength := GetMaximumMatchLength(groupSizeList)
	paddingLength := len(springs) - maximumMatchLength
	if paddingLength < 0 {
		return nil, fmt.Errorf("total groups size (%v) must not exceed report length (%v)", maximumMatchLength, len(springs))
	}
	groupsMatchStringList := make([]string, 0, len(groupSizeList))
	groupsMatchStringList, err = GetMatchStringListForGroups(groupSizeList, separatorMatchString)
	if err != nil {
		return nil, err
	}
	reportMatchStringList, err = GetAllReportMatchStrings(groupsMatchStringList, paddingLength, separatorMatchString)

	return
}
