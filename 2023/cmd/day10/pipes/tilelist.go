package pipes

import (
	"fmt"
	"strings"
)

type TileList map[Location]*Tile

/*
**	Constructors
 */

func NewTileList(capacity int) *TileList {
	tl := make(TileList, capacity)
	return &tl
}

func (tl *TileList) Add(t *Tile) {
	(*tl)[t.Location] = t
}

func (tl TileList) String() string {
	result := new(strings.Builder)
	result.WriteString(fmt.Sprintf("{Length:%v,[", len(tl)))
	for l, tile := range tl {
		result.WriteString(fmt.Sprintf("{%v: %v},", l, tile))
	}
	result.WriteString("]}")
	return result.String()
}
