package pipes

import "strings"

type PipeType string

const (
	NORTH_SOUTH_PIPE PipeType = "|"
	NORTH_EAST_PIPE  PipeType = "L"
	NORTH_WEST_PIPE  PipeType = "J"
	SOUTH_EAST_PIPE  PipeType = "F"
	SOUTH_NORTH_PIPE          = NORTH_SOUTH_PIPE
	SOUTH_WEST_PIPE  PipeType = "7"
	EAST_NORTH_PIPE           = NORTH_EAST_PIPE
	EAST_SOUTH_PIPE           = SOUTH_EAST_PIPE
	EAST_WEST_PIPE   PipeType = "-"
	WEST_EAST_PIPE            = EAST_WEST_PIPE
	WEST_NORTH_PIPE           = NORTH_WEST_PIPE
	WEST_SOUTH_PIPE           = SOUTH_WEST_PIPE
	START_PIPE       PipeType = "S"
	NO_PIPE          PipeType = "."
)

func (pt PipeType) CanConnectTo(d Direction) bool {
	return map[PipeType]Directions{
		NORTH_SOUTH_PIPE: {DIR_NORTH: true, DIR_SOUTH: true},
		NORTH_EAST_PIPE:  {DIR_NORTH: true, DIR_EAST: true},
		NORTH_WEST_PIPE:  {DIR_NORTH: true, DIR_WEST: true},
		SOUTH_EAST_PIPE:  {DIR_SOUTH: true, DIR_EAST: true},
		SOUTH_WEST_PIPE:  {DIR_SOUTH: true, DIR_WEST: true},
		EAST_WEST_PIPE:   {DIR_EAST: true, DIR_WEST: true},
		START_PIPE:       {DIR_NORTH: true, DIR_SOUTH: true, DIR_EAST: true, DIR_WEST: true},
		NO_PIPE:          {},
	}[pt][d]
}

func GetPipeTypeFromLabel(pt string) PipeType {
	switch pt {
	case "NORTH_SOUTH_PIPE":
		return NORTH_SOUTH_PIPE
	case "NORTH_EAST_PIPE":
		return NORTH_EAST_PIPE
	case "NORTH_WEST_PIPE":
		return NORTH_WEST_PIPE
	case "SOUTH_EAST_PIPE":
		return SOUTH_EAST_PIPE
	case "SOUTH_WEST_PIPE":
		return SOUTH_WEST_PIPE
	case "EAST_WEST_PIPE":
		return EAST_WEST_PIPE
	case "START_PIPE":
		return START_PIPE
	case "NO_PIPE":
		return NO_PIPE
	}
	return NO_PIPE
}

func (t PipeType) String() string {
	return string(t)
}

func (t PipeType) ToDirections() []Direction {
	if t == START_PIPE || t == NO_PIPE {
		return nil
	}
	directions := make([]Direction, 0, 2)
	ptLabel := t.ToLabel()
	dLabels := strings.Split(ptLabel, "_")
	for _, dLabel := range dLabels[0 : len(dLabels)-1] { //lop off "PIPE"
		directions = append(directions, GetDirectionFromLabel(dLabel))
	}
	return directions
}

func (t PipeType) ToLabel() string {
	switch t {
	case NORTH_SOUTH_PIPE:
		return "NORTH_SOUTH_PIPE"
	case NORTH_EAST_PIPE:
		return "NORTH_EAST_PIPE"
	case NORTH_WEST_PIPE:
		return "NORTH_WEST_PIPE"
	case SOUTH_EAST_PIPE:
		return "SOUTH_EAST_PIPE"
	case SOUTH_WEST_PIPE:
		return "SOUTH_WEST_PIPE"
	case EAST_WEST_PIPE:
		return "EAST_WEST_PIPE"
	case START_PIPE:
		return "START_PIPE"
	case NO_PIPE:
		return "NO_PIPE"
	}
	return string(t)
}
