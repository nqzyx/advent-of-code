package utils

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

func Abs[T Numeric](value T) T {
	if value >= T(0) {
		return value
	}
	if valueType := fmt.Sprintf("%T", value); strings.HasPrefix(valueType, "uint") {
		return value
	}
	return 0 - value
}

func Min[T constraints.Ordered](values ...T) T {
	var minimum T = values[0]
	for _, v := range values[1:] {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}

func Max[T constraints.Ordered](values ...T) T {
	var maximum T = values[0]
	for _, v := range values[1:] {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}

func Sum[T Numeric](values ...T) T {
	sum := values[0]
	for _, v := range values[1:] {
		sum += T(v)
	}
	return sum
}

func SumOfIntegers[T constraints.Integer](a T) T {
	var l T = T(1)
	var n T = a
	return n * (a + l) / T(2)
}
