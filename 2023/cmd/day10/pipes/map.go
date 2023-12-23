package pipes

import (
	"fmt"
	"math"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/utils"
)

type Map struct {
	colCount  int
	rowCount  int
	Insiders  *TileList
	Path      *TileList
	StartTile *Tile
	Tiles     []*TileList
}

func NewMap(inputRows *[]string) (*Map, error) {
	var err error
	// func NewTileGrid(rowCapacity, colCapacity int) *TileGrid {
	g := new(Map)
	g.rowCount = len(*inputRows)
	g.colCount = len((*inputRows)[0])
	g.Tiles = make([]*TileList, 0, g.rowCount)

	for r, inputRow := range *inputRows {
		if inputRow = strings.TrimSpace(inputRow); len(inputRow) == 0 {
			continue
		}

		if g.colCount != len(inputRow) {
			return nil, fmt.Errorf("all rows must have the same number of columns. Expected %v, found %v", g.colCount, len(inputRow))
		}

		theRow := NewTileList(len(inputRow))
		for c, inputCol := range inputRow {
			pipeType := PipeType(inputCol)
			theTile := NewTile(r, c, pipeType)
			if pipeType == StarterPipe {
				theTile.OnPath = true
				g.StartTile = theTile
			}
			*theRow = append(*theRow, *theTile)
		}
		g.Tiles = append(g.Tiles, theRow)
	}

	if g.StartTile == nil {
		return nil, fmt.Errorf("a map must have a start tile")
	}

	if g.Path, err = g.FindPath(*g.StartTile); err != nil {
		return nil, err
	}

	if g.Insiders, err = g.FindInsiders(); err != nil {
		return nil, err
	}
	return g, nil
}

func (m Map) Col(col int) *TileList {
	tileList := new(TileList)
	*tileList = make(TileList, 0, len(*m.Tiles[0]))
	for _, row := range m.Tiles {
		if col > len(*row) {
			panic(fmt.Errorf("%v exceeds the number of columns (%v)", col, len(*row)))
		}
		*tileList = append(*tileList, (*row)[col])
	}
	return tileList
}

func (m Map) Diagonal(tile Tile, rowDirection int, colDirection int) (*TileList, error) {
	if rowDirection != 1 && rowDirection != -1 && colDirection != 1 && colDirection != -1 {
		return nil, fmt.Errorf("rowDirection and colDirection must be 1 or -1")
	}
	tileList := new(TileList)
	tRow := tile.Row()
	tCol := tile.Col()
	startOffset := int(math.Abs(math.Min(float64(tRow), float64(tCol))))
	startRow := tRow - (rowDirection * startOffset)
	startCol := tCol - (colDirection * startOffset)
	var endRow int
	if rowDirection == -1 {
		endRow = 0
	} else {
		endRow = m.RowCount()
	}
	var endCol int
	if colDirection == -1 {
		endCol = 0
	} else {
		endCol = m.ColCount()
	}
	for row, col := startRow, startCol; row <= endRow && col <= endCol; row, col = row+1, col+1 {
		*tileList = append(*tileList, *m.TileAt(row, col))
	}
	return tileList, nil
}

func (m Map) TopRightDiagonal(tile Tile) (*TileList, error) {
	if tl, err := m.Diagonal(tile, +1, -1); err != nil {
		return nil, err
	} else {
		return tl, nil
	}
}

func (m Map) TopLeftDiagonal(tile Tile) (*TileList, error) {
	if tl, err := m.Diagonal(tile, +1, +1); err != nil {
		return nil, err
	} else {
		return tl, nil
	}
}

func (m Map) ColCount() int {
	return len(*(m.Tiles)[0])
}

func (m *Map) FindInsiders() (*TileList, error) {
	if m == nil {
		return nil, fmt.Errorf("tile grid does not exist")
	}
	if len(m.Tiles) == 0 {
		return nil, fmt.Errorf("tile grid is empty")
	}

	insiders := NewTileList(len(m.Tiles) / 3)

	for r := 1; r < len(m.Tiles)-2; r++ {
		row := m.Tiles[r]
		for c := 1; c < len(*row)-2; c++ {
			tile := m.TileAt(r, c)
			if tile.OnPath {
				continue
			}
			col := m.Col(c)
			fmt.Printf("r: %v, c: %v, row: %v\n", r, c, TileList((*row)[0:c]))
			leftRowTiles := TileList(utils.Reverse(TileList((*row)[0:c])))
			rightRowTiles := TileList((*row)[c+1:])
			topColTiles := TileList(utils.Reverse(TileList((*col)[0:r])))
			btmColTiles := TileList((*col)[r+1:])
			if leftRowTiles.HasPathTiles() && rightRowTiles.HasPathTiles() && topColTiles.HasPathTiles() && btmColTiles.HasPathTiles() &&
				(leftRowTiles.IsInsidePath(West) || rightRowTiles.IsInsidePath(East) || topColTiles.IsInsidePath(North) || btmColTiles.IsInsidePath(South)) {
				tile.SetPipeType(InsidePath)
				*insiders = append(*insiders, *tile)
			}
		}
	}
	return insiders, nil
}

func (m *Map) FindPath(start Tile) (*TileList, error) {
	if start.PipeType != StarterPipe {
		return nil, fmt.Errorf("start tile must have PipeType == StarterPipe")
	}

	path := NewTileList(m.RowCount() * m.ColCount() / 3)
	*path = append(*path, start)

	currentDirection, currentTile := Unknown, &start

	for {
		connectedTiles := m.GetConnectedTiles(currentDirection, currentTile)
		if connectedTiles == nil {
			return nil, fmt.Errorf("%v has no connected tiles", currentTile)
		} else if connectedTileCount := len(*connectedTiles); connectedTileCount < 1 || connectedTileCount > 2 {
			return nil, fmt.Errorf("a tile should have 1 or 2 connecting tiles, but %v has %v", currentTile, connectedTileCount)
		} else {
			for nextDirection, nextTile := range *connectedTiles {
				currentDirection = nextDirection
				nextTile.OnPath = true
				currentTile = nextTile
				break
			}
			*path = append(*path, *currentTile)
			if currentTile.PipeType == StarterPipe {
				return path, nil
			}
		}
	}
}

func (m *Map) GetConnectedTiles(from Direction, t *Tile) *map[Direction]*Tile {
	connectedTiles := make(map[Direction]*Tile)

	if from != North && t.ConnectsTo(North) && t.Row() > 0 {
		if nt := m.TileAt(t.Row()-1, t.Col()); nt.ConnectsTo(South) {
			connectedTiles[South] = nt
		}
	}
	if from != South && t.ConnectsTo(South) && t.Row() < m.RowCount()-1 {
		if nt := m.TileAt(t.Row()+1, t.Col()); nt.ConnectsTo(North) {
			connectedTiles[North] = nt
		}
	}
	if from != East && t.ConnectsTo(East) && t.Col() < m.ColCount()-1 {
		if nt := m.TileAt(t.Row(), t.Col()+1); nt.ConnectsTo(West) {
			connectedTiles[West] = nt
		}
	}
	if from != West && t.ConnectsTo(West) && t.Col() > 0 {
		if nt := m.TileAt(t.Row(), t.Col()-1); nt.ConnectsTo(East) {
			connectedTiles[East] = nt
		}
	}

	if len(connectedTiles) > 0 {
		return &connectedTiles
	} else {
		return nil
	}
}

func (m Map) Row(row int) *TileList {
	tileList := new(TileList)
	*tileList = make(TileList, 0, len(*m.Tiles[row]))
	copy(*tileList, *m.Tiles[row])
	return m.Tiles[row]
}

func (m Map) RowCount() int {
	return len(m.Tiles)
}

func (m *Map) PathLength() int {
	return len(*m.Path) / 2
}

func (m *Map) SetTile(row, col int, t *Tile) {
	(*m.Tiles[row])[col] = *t
}

func (m *Map) TileAt(row, col int) *Tile {
	return &(*m.Tiles[row])[col]
}
