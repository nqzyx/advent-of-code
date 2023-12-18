package utils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Parse[T constraints.Integer](s string, v *T) (err error) {
	_, err = fmt.Sscanf(s, "%v", v)
	return
}
