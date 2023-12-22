package pipes

type PipeType string

const (
	NorthSouth  PipeType = "|"
	EastWest    PipeType = "-"
	NorthEast   PipeType = "L"
	SouthEast   PipeType = "F"
	SouthWest   PipeType = "7"
	NorthWest   PipeType = "J"
	StarterPipe PipeType = "S"
	NoPipe      PipeType = "."
)

type (
	Connections map[PipeType]Directions
)

func (pt PipeType) ConnectsTo(dir Direction) bool {
	connections := map[PipeType]Directions{
		EastWest:    {East: true, West: true},
		NorthEast:   {North: true, East: true},
		NorthSouth:  {North: true, South: true},
		NorthWest:   {North: true, West: true},
		SouthEast:   {South: true, East: true},
		SouthWest:   {South: true, West: true},
		StarterPipe: {North: true, South: true, East: true, West: true},
		NoPipe:      {},
	}
	return connections[pt][dir]
}
