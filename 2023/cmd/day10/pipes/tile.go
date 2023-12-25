package pipes

import (
	"github.com/paulmach/orb"
)

type (
	Location  orb.Point
	Neighbors map[Direction]*Tile

	Tile/* MAYBE */ struct {
		Location  Location
		PipeType  PipeType
		OnPath    bool
		Neighbors Neighbors
	}
)

func NewTile(row, col int, pipeType PipeType, onPath bool) *Tile {
	return &Tile{
		Location: Location{float64(row), float64(col)},
		PipeType: pipeType,
		OnPath:   onPath,
	}
}

func (t Tile) Col() int {
	return int(t.Location[1])
}

func (t Tile) Row() int {
	return int(t.Location[0])
}

func (t Tile) CanConnectTo(d Direction) bool {
	if t.Neighbors != nil {
		if _, ok := t.Neighbors[d]; ok {
			return true
		}
	}
	return t.PipeType.CanConnectTo(d)
}
