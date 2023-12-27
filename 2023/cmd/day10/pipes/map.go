package pipes

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type Map struct {
	colCount  int
	rowCount  int
	Insiders  Insiders
	Path      Path
	StartTile *Tile
	Tiles     TileList
}

func NewMap(input *[]string) (m *Map, err error) {
	inputRows := slices.DeleteFunc(*input, func(inputRow string) bool {
		return len(strings.TrimSpace(inputRow)) == 0
	})

	m = new(Map)
	m.rowCount = len(inputRows)
	m.colCount = len(inputRows[0])
	var startTile *Tile
	// Add tiles to Tiles
	tiles := NewTileList(m.rowCount*m.colCount)
	for r, inputRow := range inputRows {
		if len(inputRows[0]) != len(inputRow) {
			return nil, fmt.Errorf("every row must have the same number of columns. Expected %v, found %v", len(inputRows[0]), len(inputRow))
		}
		for c, inputCol := range inputRow {
			pipeType := PipeType(inputCol)
			theTile := NewTile(r, c, pipeType, pipeType == PT_START)
			if pipeType == PT_START {
				if startTile != nil {
					return nil, fmt.Errorf("the map must not have multiple start tiles (%v, %v)", startTile, theTile.Location)
				}
				startTile = theTile
			}
			tiles.Add(theTile)
		}
	}
	m.Tiles = *tiles

	if startTile == nil {
		return nil, fmt.Errorf("the map must have a start tile")
	}

	m.StartTile = startTile

	if path, err := m.FindPath(m.StartTile); err != nil {
		return nil, err
	} else {
		m.Path = *path
	}

	if insiders, err := NewInsiders(m.Path, m.Tiles); err != nil {
		return nil, err
	} else {
		m.Insiders = *insiders
	}
	return m, nil
}

func (m *Map) isValid() bool {
	return m != nil
}

func (m *Map) ColCount() int {
	if !m.isValid() {
		return 0
	}
	return m.colCount
}

func (m *Map) RowCount() int {
	if !m.isValid() {
		return 0
	}
	return m.rowCount
}

func (m *Map) PathLength() int {
	if !m.isValid() {
		return 0
	}
	return (len(m.Path) - 1) / 2
}

func (m *Map) TileAtLocation(l Location) (*Tile, bool) {
	if ok := m.isValid(); !ok {
		return nil, ok
	}
	return m.Tiles[l], true
}

func (m *Map) TileAt(r, c int) (*Tile, bool) {
	if !m.isValid() {
		return nil, false
	}
	// check the "actual" limits
	if r < 0 || r > m.RowCount()-1 || c < 0 || c > m.ColCount()-1 {
		return nil, false
	}
	l := GetLocation(r, c)
	// get the tile
	return m.TileAtLocation(l)
}

func (m *Map) FindPath(startTile *Tile) (*Path, error) {
	var currentTile, nextTile *Tile
	var currentDirection, nextDirection Direction = DIR_UNKNOWN, DIR_UNKNOWN
	// Ensure the start tile has been set
	currentTile = startTile
	if currentTile == nil {
		return nil, fmt.Errorf("start tile for the map must exist")
	}
	// Ensure the start tile is really the start tile
	if currentTile.PipeType != PT_START {
		return nil, fmt.Errorf("start tile must have valid pipe type (expected: \"%v\", got: \"%v\")", PT_START.String(), currentTile.PipeType.String())
	}
	// Init path
	path := NewPath()

	// iterate through the entire path
	for {
		// Add current tile to path
		path.Add(currentTile)
		// Get connected neighbors
		connectedNeighbors := m.GetConnectedNeighbors(currentTile)
		if len(*connectedNeighbors) != 2 {
			return nil, fmt.Errorf("each tile on the path should have 2 connected neighbors, but %v has %v", currentTile, len(*connectedNeighbors))
		}
		// Add Neighbors to current tile
		if currentTile.Neighbors == nil {
			currentTile.Neighbors = make(Neighbors, 0)
		}
		// transfer connectedNeighbors to currentTile
		for d, t := range *connectedNeighbors {
			currentTile.Neighbors[d] = Neighbor(t)
			if d != currentDirection.Opposite() || currentDirection == DIR_UNKNOWN {
				nextTile = t
				nextDirection = d
				nextTile.OnPath = currentTile.OnPath
			}
		}
		// set pipe type for the start tile; requires neighbors to be defined
		if currentTile.PipeType == PT_START {
			// set the start tile's pipe type based on directions to neighbors
			directionNames := make([]string, 0, len(currentTile.Neighbors))
			for _, d := range maps.Keys(currentTile.Neighbors) {
				directionNames = append(directionNames, d.String())
			}
			currentTile.PipeType = GetPipeType(strings.Join(directionNames, "_"))
		}

		sr, sc := startTile.Location.GetRowCol()
		nr, nc := nextTile.Location.GetRowCol()
		if sr == nr && sc == nc {
			path.Add(nextTile)
			break
		}

		if path.Closed() {
			return nil, fmt.Errorf("the path must start and end with the start tile (%v)", startTile)
		}

		// Set up for next tile on the path
		currentTile = nextTile
		currentDirection = nextDirection
	}

	return path, nil
}

func (m *Map) GetNeighbor(t *Tile, d Direction) (*Tile, bool) {
	r, c := t.Location.GetRowCol()
	switch d {
	case DIR_NORTH:
		r--
	case DIR_SOUTH:
		r++
	case DIR_EAST:
		c++
	case DIR_WEST:
		c--
	default:
		return nil, false
	}
	if n, ok := (*m).TileAt(r, c); !ok {
		return nil, ok
	} else {
		return n, true
	}
}

func (m *Map) GetConnectedNeighbors(t *Tile) *map[Direction]*Tile {
	nArr := make(map[Direction]*Tile)
	dArr := t.PipeType.ToDirections()

	for _, d := range dArr {
		if n, ok := m.GetNeighbor(t, d); !ok || n == nil {
			continue
		} else if n.CanConnect(d.Opposite()) {
			nArr[d] = n
		}
	}
	return &nArr
}
