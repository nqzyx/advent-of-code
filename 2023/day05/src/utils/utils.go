package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

func PrintJSON(v any) string {
	indent := strings.Repeat(" ", 2)
	ba, err := json.MarshalIndent(v, "", indent)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(string(ba))
}

var integerRegexp = regexp.MustCompile("[[:digit:]]+")

func Parse[T constraints.Integer](s string, v *T) (err error) {
	_, err = fmt.Sscanf(s, "%v", v)
	return
}

func NewIntArrayFromString[T constraints.Integer](s string) (result []T) {
	s = strings.TrimSpace(s)
	sArr := integerRegexp.FindAllString(s, -1)
	result = make([]T, len(sArr))
	for i, v := range sArr {
		var x T
		err := Parse(v, &x)
		if err != nil {
			panic(err)
		}
		result[i] = T(x)
	}
	return
}
