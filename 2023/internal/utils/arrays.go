package utils

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Float | constraints.Integer
}

func NewIntArrayComma[T constraints.Integer, S ~string](s S) []T {
	return NewIntArray[T](string(s), ",")
}

func NewIntArraySpace[T constraints.Integer, S ~string](s S) []T {
	return NewIntArray[T](string(s), " ")
}

func NewIntArray[T constraints.Integer, S ~string](s S, sep S) []T {
	sArr := strings.Split(strings.TrimSpace(string(s)), string(sep))
	result := make([]T, 0, len(sArr))
	for _, s := range sArr {
		var x T
		if l, err := fmt.Sscanf(s, "%v", &x); err != nil {
			panic(err)
		} else {
			if l == 0 {
				return result
			}
			result = append(result, T(x))
		}
	}
	return result
}
