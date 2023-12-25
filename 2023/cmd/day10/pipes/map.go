package pipes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nqzyx/advent-of-code/utils"
	"github.com/paulmach/orb"
	"golang.org/x/exp/maps"
)

type (
	Insiders orb.MultiPoint

	Map struct {
		colCount  int
		rowCount  int
		Insiders  *Insiders
		Path      *Path
		StartTile *Location
		Tiles     [][]Tile
	}
)

func NewMap(inputRows *[]string) (*Map, error) {
	var err error

	slices.DeleteFunc(*inputRows, func(inputRow string) bool {
		return len(strings.TrimSpace(inputRow)) == 0
	})

	m := *new(Map)
	m.rowCount = len(*inputRows)
	m.colCount = len((*inputRows)[0])
	m.Tiles = make([][]Tile, 0, m.rowCount)

	for r, inputRow := range *inputRows {
		if m.ColCount() != len(inputRow) {
			return nil, fmt.Errorf("all rows must have the same number of columns. Expected %v, found %v", m.ColCount(), len(inputRow))
		}

		for c, inputCol := range inputRow {
			pipeType := PipeType(inputCol)
			tileLocation := Location{float64(r), float64(c)}
			theTile := *NewTile(r, c, pipeType, pipeType == START_PIPE)
			if pipeType == START_PIPE {
				m.StartTile = &tileLocation
			}
			m.AddTile(theTile)
		}
	}

	if m.StartTile == nil {
		panic(fmt.Errorf("the map must have a starting tile"))
	}

	if m.Path, err = m.FindPath(*m.StartTile); err != nil {
		return nil, err
	}

	// if m.Insiders, err = m.FindInsiders(); err != nil {
	// 	return nil, err
	// }
	return &m, nil
}

func (m *Map) AddTile(t Tile) {
	if m.Tiles == nil {
		m.Tiles = make([][]Tile, 0, m.rowCount)
	}
	row, col := GetRowCol(t.Location)
	if m.Tiles[row] == nil {
		m.Tiles[row] = make([]Tile, 0, m.colCount)
	}
	m.Tiles[row][col] = t
}

func (m *Map) TileAtLocation(l Location) *Tile {
	row, col := GetRowCol(l)
	return &m.Tiles[row][col]
}

func (m *Map) TileAt(r, c int) *Tile {
	return &m.Tiles[r][c]
}

func (m *Map) ColCount() int {
	return m.colCount
}

func (m *Map) FindInsiders() (*TileList, error) {
	utils.FunctionNotImplemented()
	// if m == nil {
	// 	return nil, fmt.Errorf("tile grid does not exist")
	// }
	// if len(m.Tiles) == 0 {
	// 	return nil, fmt.Errorf("tile grid is empty")
	// }

	// insiders := NewTileList(len(m.Tiles) / 3)

	// for r := 1; r < len(m.Tiles)-2; r++ {
	// 	row := m.Tiles[r]
	// 	for c := 1; c < len(*row)-2; c++ {
	// 		tile := m.TileAt(r, c)
	// 		if tile.OnPath {
	// 			continue
	// 		}
	// 		// col := m.Col(c)
	// 		// fmt.Printf("r: %v, c: %v, row: %v\n", r, c, TileList((*row)[0:c]))
	// 		// eastWestTiles := row
	// 		// eastTiles := (*eastWestTiles)[0:c]
	// 		// westTiles := (*eastWestTiles)[c+1:]
	// 		// northSouthTiles := m.Col(c)
	// 		// northTiles := utils.Reverse((*northSouthTiles)[0:r])
	// 		// btmColTiles := TileList((*col)[r+1:])
	// 		// if eastTiles.EnclosesTilesOn(WEST) || westTiles.EnclosesTilesOn(DIR_EAST) {
	// 		// 	tile.PipeType = InsidePath
	// 		// 	*insiders = append(*insiders, *tile)
	// 		// }
	// 	}
	// }
	// return insiders, nil
	return nil, nil
}

func (m *Map) FindPath(l Location) (*Path, error) {
	var currentTile, nextTile *Tile
	var currentDirection, nextDirection Direction = DIR_UNKNOWN, DIR_UNKNOWN
	// Ensure the start tile has been set and is valid
	currentTile = m.TileAtLocation(l)
	if currentTile == nil {
		return nil, fmt.Errorf("map has no start tile")
	}
	if currentTile.PipeType != START_PIPE {
		return nil, fmt.Errorf("invalid start tile (pipe type is \"%v\")", currentTile.PipeType)
	}
	// hang on to the start tile for later comparisons
	startTile := currentTile
	// Init path
	path := new(Path)
	path.Steps = make(orb.Ring, 0, len(m.Tiles)/3)
	path.Sides = make(orb.MultiLineString, 0)
	// iterate through the entire path
	for {
		// Add current tile to path
		path.AddTile(currentTile)
		connectedNeighbors := *m.GetConnectedNeighbors(currentTile, currentDirection)
		if len(connectedNeighbors) != 2 {
			return nil, fmt.Errorf("each tile on the path should have 2 connected neighbors, but %v has %v", currentTile, len(connectedNeighbors))
		}
		// Add Neighbors to current tile
		for d, t := range connectedNeighbors {
			currentTile.Neighbors[d] = t
			if d != currentDirection {
				nextTile = t
				nextDirection = d
				nextTile.OnPath = currentTile.OnPath
			}
		}
		// set pipe type for the start tile
		if orb.Point(currentTile.Location).Equal(orb.Point(startTile.Location)) {
			// set the start tile's pipe type based on directions to neighbors
			directionLabels := make([]string, 0, len(connectedNeighbors))
			for _, d := range maps.Keys(connectedNeighbors) {
				directionLabels = append(directionLabels, string(d.ToLabel()))
			}
			pipeTypeLabel := strings.Join(directionLabels, "_")
			currentTile.PipeType = GetPipeTypeFromLabel(pipeTypeLabel)
		}
		// coming back to start tile means we're done
		if orb.Point(nextTile.Location).Equal(orb.Point(startTile.Location)){
			break;
		}
		// Set up for next tile on the path
		currentTile = nextTile
		currentDirection = nextDirection
	}
	// TODO: set up (Path).Sides from (Path).Steps
	return path, nil
}

func (m *Map) GetNeighbor(t *Tile, d Direction) (*Tile, Direction) {
	r, c := GetRowCol(t.Location)
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
		return nil, d
	}
	return (*m).TileAt(r, c), d
}

func (m *Map) GetConnectedNeighbors(t *Tile, d Direction) *map[Direction]*Tile {
	connectedNeighbors := make(map[Direction]*Tile)
	var neighborTile *Tile
	directions := []Direction{DIR_NORTH, DIR_SOUTH, DIR_EAST, DIR_WEST}
	for _, neighborDirection := range directions {
		if neighborTile, _ = m.GetNeighbor(t, d); neighborTile != nil {
			if t.CanConnectTo(neighborDirection) && neighborTile.CanConnectTo(d.GetOppositeDirection()) {
				connectedNeighbors[neighborDirection] = neighborTile
			}
		}
	}
	return &connectedNeighbors
}

func (m Map) RowCount() int {
	return m.rowCount
}

func (m *Map) PathLength() int {
	return len(m.Path.Steps) / 2
}

