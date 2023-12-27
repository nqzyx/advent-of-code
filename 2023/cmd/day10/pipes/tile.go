package pipes

import (
	"fmt"
	"strings"
	// "github.com/paulmach/orb"
)

type (
	Neighbor *Tile

	Neighbors map[Direction]Neighbor

	Tile struct {
		Location  Location
		PipeType  PipeType
		OnPath    bool
		Neighbors Neighbors
	}
)

/*
**	Constructors
 */

// NewTile(row, col int, pipeType PipeType, onPath bool) *Tile
func NewTile(row, col int, pipeType PipeType, onPath bool) *Tile {
	return &Tile{
		Location: Location{float64(row), float64(col)},
		PipeType: pipeType,
		OnPath:   onPath,
	}
}

// (Tile).CanConnectTo(d Direction) bool
func (t Tile) CanConnect(d Direction) bool {
	if t.Neighbors != nil {
		if _, ok := t.Neighbors[d]; ok {
			return true
		}
	}
	return t.PipeType.CanConnect(d)
}

// // (Tile).Col() int
// func (t Tile) Col() int {
// 	return int(orb.Point(t.Location).X())
// }

// // (Tile).Row() int
// func (t Tile) Row() int {
// 	return int(orb.Point(t.Location).Y())
// }

// (Tile).String string
func (t Tile) String() string {
	sb := new(strings.Builder)
	sb.WriteString("{")
	sb.WriteString(
		strings.Join([]string{
			fmt.Sprintf("%v", t.Location),
			fmt.Sprintf("Type:%v", t.PipeType),
			fmt.Sprintf("Path:%v", t.OnPath),
		},
			",",
		),
	)
	sb.WriteString(",Neighbors:[")
	for d, n := range t.Neighbors {
		sb.WriteString(fmt.Sprintf("{%v:%v}", d,
			fmt.Sprintf("{%v,OnPath:%v,%v}", n.Location, n.OnPath, n.PipeType),
		))
	}
	sb.WriteString("]")
	sb.WriteString("}")
	return sb.String()
}
