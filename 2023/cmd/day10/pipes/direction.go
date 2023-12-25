package pipes

import (
	"strings"
)

type Direction string

const (
	// Cardinal Directions
	DIR_UNKNOWN Direction = "U"
	DIR_NORTH   Direction = "N"
	DIR_SOUTH   Direction = "S"
	DIR_EAST    Direction = "E"
	DIR_WEST    Direction = "W"
)

type Directions map[Direction]bool

func (d Direction) GetOppositeDirection() Direction {
	switch d {
	case DIR_NORTH:
		return DIR_SOUTH
	case DIR_SOUTH:
		return DIR_NORTH
	case DIR_EAST:
		return DIR_WEST
	case DIR_WEST:
		return DIR_EAST
	case DIR_UNKNOWN:
		return DIR_UNKNOWN
	default:
		return DIR_UNKNOWN
	}
}

func (d Direction) String() string {
	return string(d)
}

func (d Direction) ToLabel() string {
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

func (d1 Direction) ToPipeType(d2 Direction) PipeType {
	if d1 == DIR_UNKNOWN || d2 == DIR_UNKNOWN {
		return NO_PIPE
	}
	dLabels := make([]string, 0, 2)
	dLabels = append(dLabels, d1.ToLabel(), d2.ToLabel())
	ptLabel := strings.Join(dLabels, "_")
	pt := GetPipeTypeFromLabel(ptLabel)
	return pt
}

// helpers

func GetDirectionFromLabel(l string) Direction {
	switch l {
	case "DIR_NORTH":
		return DIR_NORTH
	case "DIR_SOUTH":
		return DIR_SOUTH
	case "DIR_EAST":
		return DIR_EAST
	case "DIR_WEST":
		return DIR_WEST
	case "DIR_UNKNOWN":
		return DIR_UNKNOWN
	default:
		return DIR_UNKNOWN
	}

}
