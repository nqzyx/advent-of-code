package utils

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

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
