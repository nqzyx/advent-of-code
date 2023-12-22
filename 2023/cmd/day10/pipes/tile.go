package pipes

type PathPosition string

const (
	OnPath      PathPosition = "P"
	InsidePath  PathPosition = "I"
	OutsidePath PathPosition = "O"
)

type (
	Coordinates [2]int

	Tile struct {
		Coords       Coordinates
		PipeType     PipeType
		PathPosition PathPosition
	}

	TileRow  []Tile
	TilePath []Tile
	TileGrid []TileRow
)

func NewTile(row, col int, pipeType PipeType) (t *Tile) {
	return &Tile{
		Coords:   Coordinates{row, col},
		PipeType: pipeType,
	}
}

func (t Tile) Col() int {
	return t.Coords[1]
}

func (t Tile) Row() int {
	return t.Coords[0]
}

func (t Tile) ConnectsTo(d Direction) bool {
	return t.PipeType.ConnectsTo(d)
}

func (g *TileGrid) Row(r int) TileRow {
	return (*g)[r]
}
