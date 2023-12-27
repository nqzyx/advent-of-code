package pipes

type (
	Direction        string
	Directions       []Direction
	DirectionDetails struct {
		opposite Direction
	}
)

const (
	// Cardinal Directions
	DIR_UNKNOWN Direction = "U"
	DIR_NORTH   Direction = "N"
	DIR_SOUTH   Direction = "S"
	DIR_EAST    Direction = "E"
	DIR_WEST    Direction = "W"
)

/*
**	VARIABLES
 */

var AllDirections = map[Direction]DirectionDetails{
	DIR_NORTH:   {DIR_SOUTH},
	DIR_SOUTH:   {DIR_NORTH},
	DIR_EAST:    {DIR_WEST},
	DIR_WEST:    {DIR_EAST},
	DIR_UNKNOWN: {DIR_UNKNOWN},
}

/*
**	METHODS
 */

func (d Direction) Opposite() Direction {
	return AllDirections[d].opposite
}

func (d Direction) Name() string {
	switch d {
	case DIR_NORTH:
		return "DIR_NORTH"
	case DIR_SOUTH:
		return "DIR_SOUTH"
	case DIR_EAST:
		return "DIR_EAST"
	case DIR_WEST:
		return "DIR_WEST"
	case DIR_UNKNOWN:
		return "DIR_UNKNOWN"
	default:
		return "DIR_UNKNOWN"
	}
}

func (d Direction) String() string {
	return d.Name()[len("DIR_"):]
}
