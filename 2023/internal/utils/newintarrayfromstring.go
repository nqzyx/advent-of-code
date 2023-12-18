package utils

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

var integerRegexp = regexp.MustCompile("[+-]?[[:digit:]]+")

func NewIntArrayFromString[T constraints.Integer](s string) (result []T) {
	s = strings.TrimSpace(s)
	sArr := integerRegexp.FindAllString(s, -1)
	result = make([]T, len(sArr))
	for i, v := range sArr {
		var x T
		err := Parse(v, &x)
		if err != nil {
			fmt.Println(err)
		}
		result[i] = T(x)
	}
	return
}
