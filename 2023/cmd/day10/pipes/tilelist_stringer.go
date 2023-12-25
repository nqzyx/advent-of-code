package pipes

import (
	"fmt"
	"strings"
)

func (tl TileList) String() string {
	result := new(strings.Builder)
	result.WriteString(fmt.Sprintf("Length: %v, TileList{", len(tl)))
	for i, tile := range tl {
		result.WriteString(fmt.Sprintf("%v: %v,", i, tile))
	}
	result.WriteString("}")
	return result.String()
}
