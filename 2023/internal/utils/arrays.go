package utils

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

var integerRegexp = regexp.MustCompile("[+-]?[[:digit:]]+")

func Insert[T any](a []T, c T, i int) []T {
	return append(a[:i], append([]T{c}, a[i:]...)...)
}

func Reverse[T any](a []T) []T {
	newA := make([]T, len(a))
	for i, j := 0, len(a)-1; i < len(a); i, j = i+1, j-1 {
		newA[i] = a[j]
	}
	return newA
}

type Numeric interface {
	constraints.Float | constraints.Integer
}

func NewNumericArrayFromString[T Numeric](str string) (result []T) {
	return NewNumericArrayFromStringWithSeparator[T](str, " ")
}

func NewNumericArrayFromStringWithSeparator[T Numeric](str string, sep string) (result []T) {
	sArr := strings.Split(strings.TrimSpace(str), sep)
	result = make([]T, 0, len(sArr))
	for i, s := range sArr {
		var x T
		if l, err := fmt.Sscanf(s, "%v", &x); err != nil {
			panic(err)
		} else {
			if l == 0 {
				return
			}
			result[i] = T(x)
		}
	}
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
			fmt.Println(err)
		}
		result[i] = T(x)
	}
	return
}
