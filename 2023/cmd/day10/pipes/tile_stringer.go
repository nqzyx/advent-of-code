package pipes

import (
	"fmt"
	"strings"
)

func (t Tile /* YES */) String() string {
	sb := new(strings.Builder)
	sb.WriteString("{ ")
	sb.WriteString(
		strings.Join([]string{
			fmt.Sprintf("Loc: [%v, %v]", t.Row(), t.Col()),
			fmt.Sprintf("Type: %v", t.PipeType),
			fmt.Sprintf("Path: %v", t.OnPath),
			// fmt.Sprintf("Start: %v", t.IsStarter),
		},
			", ",
		),
	)
	sb.WriteString(" }")
	return sb.String()
}
