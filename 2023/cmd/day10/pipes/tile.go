package pipes

type (
	Coordinates [2]int

	Tile struct {
		Coords         Coordinates
		PipeType       PipeType
		ConnectedTiles [2]Coordinates
	}
)

func NewTile(x, y int, pipeType PipeType) (t *Tile) {
	return &Tile{
		Coords:         Coordinates{x, y},
		PipeType:       pipeType,
		ConnectedTiles: *new([2]Coordinates),
	}
}

func (t Tile) Col() int {
	return t.Coords[0]
}

func (t Tile) Row() int {
	return t.Coords[1]
}

func (t Tile) ConnectsTo(d Direction) bool {
	return t.PipeType.ConnectsTo(d)
}
