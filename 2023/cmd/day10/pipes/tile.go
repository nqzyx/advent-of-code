package pipes

type (
	Coordinates [2]int
	Tile        struct {
		Coords Coordinates
		Pipe   Pipe
	}
)

func NewTile(x, y int, pipe Pipe) (t *Tile) {
	return &Tile{
		Coords: Coordinates{x, y},
		Pipe:   pipe,
	}
}

func (t Tile) X() int {
	return t.Coords[0]
}

func (t Tile) Y() int {
	return t.Coords[1]
}
