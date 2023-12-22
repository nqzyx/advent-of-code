package pipes

import (
	"fmt"
	"strings"
)

type Map struct {
	Rows      int
	Cols      int
	StartTile Tile
	Tiles     [][]Tile
	Path      []Tile
}

func NewMap(mapRows []string) (m *Map, err error) {
	startTileFound := false
	m = new(Map)
	m.Tiles = make([][]Tile, 0, len(mapRows))
	m.Cols = len(mapRows[0])
	for r, mapRow := range mapRows {
		if mapRow = strings.TrimSpace(mapRow); len(mapRow) == 0 {
			continue
		}
		if m.Cols != len(mapRow) {
			return nil, fmt.Errorf("all rows must have the same number of columns. Expected %v, found %v", m.Cols, len(mapRow))
		}
		row := make([]Tile, 0, len(mapRow))
		for c, colData := range mapRow {
			tile := *NewTile(c, r, PipeType(colData))
			row = append(row, tile)
			if tile.PipeType == StarterPipe {
				m.StartTile = tile
				startTileFound = true
			}
		}
		m.Tiles = append(m.Tiles, row)
	}

	if !startTileFound {
		return nil, fmt.Errorf("no start tile found for map")
	}

	m.Rows = len(m.Tiles)

	return m, nil
}

func (m *Map) GetConnectingTiles(incomingDirection Direction, tile Tile) (map[Direction]Tile, error) {
	neighbors := make(map[Direction]Tile)
	switch true {
	case incomingDirection != North && tile.ConnectsTo(North) && tile.Row() > 0:
		if neighbor := m.Tiles[tile.Row()-1][tile.Col()]; neighbor.ConnectsTo(South) {
			neighbors[South] = neighbor
		}
	case incomingDirection != South && tile.ConnectsTo(South) && tile.Row() < m.Rows-1:
		if neighbor := m.Tiles[tile.Row()+1][tile.Col()]; neighbor.ConnectsTo(North) {
			neighbors[North] = neighbor
		}
	case incomingDirection != East && tile.ConnectsTo(East) && tile.Col() < m.Cols-1:
		if neighbor := m.Tiles[tile.Row()][tile.Col()+1]; neighbor.ConnectsTo(West) {
			neighbors[West] = neighbor
		}
	case incomingDirection != West && tile.ConnectsTo(West) && tile.Col() > 0:
		if neighbor := m.Tiles[tile.Row()][tile.Col()-1]; neighbor.ConnectsTo(East) {
			neighbors[East] = neighbor
		}
	}

	delete(neighbors, incomingDirection)

	switch len(neighbors) {
	case 0:
		return nil, fmt.Errorf("tile %v does not have any connecting tiles", tile)
	case 1, 2:
		return neighbors, nil
	default:
		return nil, fmt.Errorf("a tile must only have no more than two (2) connecting tiles, not %v", len(neighbors))
	}
}

func (m *Map) FindPathLength() (int, error) {
	m.Path = make([]Tile, 0, m.Rows*m.Cols/2)
	m.Path = append(m.Path, m.StartTile)

	var tile Tile
	var currentDirection Direction

	if neighbors, err := m.GetConnectingTiles(Unknown, m.StartTile); err != nil {
		return 0, err
	} else if neighbors == nil || len(neighbors) != 2 {
		return 0, fmt.Errorf("the starting tile must have two (2) connected tiles, not %v", len(neighbors))
	} else {
		for d, t := range neighbors {
			currentDirection = d
			tile = t
			break
		}
	}

	for {
		if neighbors, err := m.GetConnectingTiles(currentDirection, tile); err != nil {
			return 0, fmt.Errorf("an error occurred when attempting to find connecting tiles at %v, (%w)", tile, err)
		} else if neighbors == nil || len(neighbors) != 1 {
			return 0, fmt.Errorf("tile (%v) should have exactly one (1) connecting tile, but has %v", tile, len(neighbors))
		} else {
			for d, t := range neighbors {
				currentDirection = d
				tile = t
				break
			}
			m.Path = append(m.Path, tile)
			if tile == m.StartTile {
				return len(m.Path) / 2, nil
			}
		}
	}
}
