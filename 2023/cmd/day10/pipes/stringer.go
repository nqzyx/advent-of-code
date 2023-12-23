package pipes

import (
	"fmt"
	"strings"
)

var pageHeaders []string = []string{
	"0....0....1....1....2....2....3....3....4....4....5....5....6....6....7....7....8....8....9....9....",
	"0....5....0....5....0....5....0....5....0....5....0....5....0....5....0....5....0....5....0....5....",
}

var (
	hdrPrefix string = "   : "
	rowPrefix string = "%-3v: "
)

func (m Map) String() string {
	lineLength := len(*m.Tiles[0])
	result := new(strings.Builder)
	for _, pageHeader := range pageHeaders {
		result.WriteString(fmt.Sprintf("%v%v\n", hdrPrefix, strings.Repeat(pageHeader, (lineLength/100)+1)[0:lineLength]))
	}

	for r, tileRow := range m.Tiles {
		switch r % 5 {
		case 0:
			result.WriteString(fmt.Sprintf(rowPrefix, r%1000))
		default:
			result.WriteString(hdrPrefix)
		}
		for _, tile := range *tileRow {
			if tile.OnPath {
				result.WriteString("~")
			} else {
				result.WriteString(string(tile.PipeType))
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func (l TileList) String() string {
	result := new(strings.Builder)
	result.WriteString(fmt.Sprintf("len: %v, TileList{", len(l)))
	for i, t := range l {
		result.WriteString(fmt.Sprintf("%v: %v,", i, t))
	}
	result.WriteString("}")
	return result.String()
}
