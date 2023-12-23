package pipes

type (
	Coordinates [2]int

	Tile struct {
		Coords   Coordinates
		PipeType PipeType
		OnPath   bool
	}
)

func NewTile(row, col int, pipeType PipeType) *Tile {
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

func (t *Tile) SetPipeType(p PipeType) (*Tile, PipeType) {
	if t.OnPath {
		switch p {
		case NoPipe:
			p = InsidePath
		case InsidePath:
			p = NoPipe
		}
	} else {
		t.PipeType = p
	}
	return t, p
}
