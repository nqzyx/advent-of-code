package utils

import (
	"fmt"
)

func Parse[T Numeric](s string, v *T) (err error) {
	_, err = fmt.Sscanf(s, "%v", v)
	return
}
