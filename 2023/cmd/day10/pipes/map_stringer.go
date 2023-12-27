package pipes

import (
	"fmt"
	"strings"
)

func (m Map) String() string {
	sb := new(strings.Builder)
	sb.WriteString(
		strings.Join([]string{
			fmt.Sprintf("RowCount:%v,ColCount:%v", m.rowCount, m.colCount),
			fmt.Sprintf("Tiles:%v", m.Tiles),
			fmt.Sprintf("Path:[%v]", m.Path),
			fmt.Sprintf("Insiders:%v", m.Insiders),
		}, ",",
		),
	)
	return sb.String()
}
