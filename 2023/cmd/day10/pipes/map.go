package pipes

import (
	"fmt"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/utils"
)

type Map struct {
	RowCount int
	ColCount int
	Grid     TileGrid
	Path     TilePath
	Insiders []Tile
}

func NewMap(inputRows []string) (m *Map, err error) {
	m = new(Map)
	m.ColCount = len(inputRows[0])
	m.Grid = make(TileGrid, 0, len(inputRows))
	m.Path = make(TilePath, 0, len(inputRows))
	for iRow, inputRow := range inputRows {
		if inputRow = strings.TrimSpace(inputRow); len(inputRow) == 0 {
			continue
		}
		if m.ColCount != len(inputRow) {
			return nil, fmt.Errorf("all rows must have the same number of columns. Expected %v, found %v", m.ColCount, len(inputRow))
		}
		theRow := make(TileRow, 0, len(inputRow))
		for iCol, inputCol := range inputRow {
			theTile := *NewTile(iRow, iCol, PipeType(inputCol))
			if PipeType(inputCol) == StarterPipe {
				theTile.PathPosition = OnPath
				m.Path = append(m.Path, theTile)
			}
			theRow = append(theRow, theTile)
		}
		m.Grid = append(m.Grid, theRow)
	}

	if starterTilesCount := len(m.Path); starterTilesCount != 1 {
		return nil, fmt.Errorf("a map must have 1 start tile, not %v", starterTilesCount)
	}

	m.RowCount = len(m.Grid)

	if err := m.FindPath(); err != nil {
		return nil, err
	}

	// utils.PrintlnJSON(m, false)

	return m, nil
}

func (m *Map) DistanceToFarthestTile() int {
	if m == nil || m.Path == nil {
		return 0
	}
	return len(m.Path) / 2
}

func (m *Map) GetConnectedTiles(entryDirection Direction, currentTile Tile) map[Direction]Tile {
	neighbors := make(map[Direction]Tile)

	if entryDirection != North && currentTile.ConnectsTo(North) && currentTile.Row() > 0 {
		if neighbor := m.Grid[currentTile.Row()-1][currentTile.Col()]; neighbor.ConnectsTo(South) {
			neighbors[South] = neighbor
		}
	}
	if entryDirection != South && currentTile.ConnectsTo(South) && currentTile.Row() < m.RowCount-1 {
		if neighbor := m.Grid[currentTile.Row()+1][currentTile.Col()]; neighbor.ConnectsTo(North) {
			neighbors[North] = neighbor
		}
	}
	if entryDirection != East && currentTile.ConnectsTo(East) && currentTile.Col() < m.ColCount-1 {
		if neighbor := m.Grid[currentTile.Row()][currentTile.Col()+1]; neighbor.ConnectsTo(West) {
			neighbors[West] = neighbor
		}
	}
	if entryDirection != West && currentTile.ConnectsTo(West) && currentTile.Col() > 0 {
		if neighbor := m.Grid[currentTile.Row()][currentTile.Col()-1]; neighbor.ConnectsTo(East) {
			neighbors[East] = neighbor
		}
	}

	if len(neighbors) > 0 {
		return neighbors
	} else {
		return nil
	}
}

func (m *Map) FindPath() error {
	if m == nil || m.Path == nil || len(m.Path) == 0 {
		return nil
	}

	if len(m.Path) > 1 {
		return fmt.Errorf("a map must have exactly one starting tile, not %v", len(m.Path))
	}

	currentDirection, currentTile := Unknown, &m.Path[0]

	for {
		neighbors := m.GetConnectedTiles(currentDirection, *currentTile)
		if neighbors == nil || len(neighbors) < 1 || len(neighbors) > 2 {
			return fmt.Errorf("a tile should have 1 or 2 connecting tiles, but %v has %v", currentTile, len(neighbors))
		} else {
			for nextDirection, nextTile := range neighbors {
				currentDirection = nextDirection
				currentTile = &nextTile
				currentTile.PathPosition = OnPath
				m.Grid[currentTile.Row()][currentTile.Col()].PathPosition = OnPath
				break
			}
			m.Path = append(m.Path, *currentTile)
			if currentTile.PipeType == StarterPipe {
				return nil
			}
		}
	}
}

func (m *Map) FindTilesInsidePath() (int, error) {
	if m == nil {
		return 0, fmt.Errorf("called on an nil pointer")
	}
	if m.Grid == nil {
		return 0, fmt.Errorf("map has no tile grid")
	}
	if len(m.Grid) == 0 {
		return 0, fmt.Errorf("tile grid is empty")
	}

	m.Insiders = make([]Tile, 0, len(m.Grid)/3)

	pathPosition := OutsidePath

	for iRow, row := range m.Grid {
		for iCol, tile := range row {
			switch tile.PathPosition {
			case OnPath:
				switch pathPosition {
				case OutsidePath:
					pathPosition = InsidePath
				case InsidePath:
					pathPosition = OutsidePath
				}
			default:
				fmt.Printf("updating %d,%d to %v\n", iRow, iCol, pathPosition)
				tile.PathPosition = pathPosition
				m.Grid[tile.Row()][tile.Col()].PathPosition = pathPosition
				if pathPosition == InsidePath {
					m.Insiders = append(m.Insiders, tile)
				}
			}
			utils.PrintlnJSON(tile, false)
		}
	}
	return len(m.Insiders), nil
}
