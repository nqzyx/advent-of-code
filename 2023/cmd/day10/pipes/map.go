package pipes

import "strings"

type Map struct {
	Start Tile
	Tiles [][]Tile
}

func NewMap(input []string) (m *Map) {
	m = new(Map)
	for r, line := range input {
		sc := strings.Index(line, string(Start))
		if sc >= 0 {
			m.Start = *NewTile(sc, r, Start)
		}
		row := make([]Tile, 0, len(line))
		for c, pipe := range line {
			row = append(row, *NewTile(c, r, Pipe(pipe)))
		}
		m.Tiles = append(m.Tiles, row)
	}
	return
}

func (m *Map) ConnectingTiles(t Tile) (adjacentTiles []Tile) {
	adjacentTiles = make([]Tile, 0, 4)
	if t.X() > 0 {
		westTile := m.Tiles[t.Y()][t.X()-1]
		if pipeConnections[westTile.Pipe].East {
			adjacentTiles = append(adjacentTiles, westTile)
		}
	}
	if t.X() < len(m.Tiles[t.Y()]) {
		eastTile := m.Tiles[t.Y()][t.X()+1]
		if pipeConnections[eastTile.Pipe].West {
			adjacentTiles = append(adjacentTiles, eastTile)
		}
	}
	if t.Y() > 0 {
		northTile := m.Tiles[t.Y()-1][t.X()]
		if pipeConnections[northTile.Pipe].South {
			adjacentTiles = append(adjacentTiles, northTile)
		}
	}
	if t.Y() < len(m.Tiles) {
		southTile := m.Tiles[t.Y()+1][t.X()]
		if pipeConnections[southTile.Pipe].North {
			adjacentTiles = append(adjacentTiles, southTile)
		}
	}
	return
}

func (m *Map) GetPathLength() (l int) {
	currentStep := m.ConnectingTiles(m.Start)[0]
	for {
		m.ConnectingTiles(currentStep)
	}

	return
}
