package utils

import (
	"golang.org/x/exp/constraints"
)

type SignedNumber interface {
	constraints.Signed | constraints.Float
}

func Min[T constraints.Ordered](values ...T) T {
	var min T = values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[T constraints.Ordered](values ...T) T {
	var max T = values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Abs[T SignedNumber](value T) T {
	if value >= T(0) {
		return value
	} else {
		return value * T(-1)
	}
}
