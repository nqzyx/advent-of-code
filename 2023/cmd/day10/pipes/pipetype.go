package pipes

import (
	"slices"

	"golang.org/x/exp/maps"
)

type (
	PipeType        string
	PipeTypeDetails struct {
		name       string
		directions Directions
	}
)

const (
	PT_NORTH_SOUTH PipeType = "|"
	PT_NORTH_EAST  PipeType = "L"
	PT_NORTH_WEST  PipeType = "J"
	PT_SOUTH_EAST  PipeType = "F"
	PT_SOUTH_WEST  PipeType = "7"
	PT_EAST_WEST   PipeType = "-"
	PT_START       PipeType = "S"
	PT_NONE        PipeType = "."
)

var allPipeTypes = map[PipeType]PipeTypeDetails{
	PT_NORTH_SOUTH: {"PT_NORTH_SOUTH", Directions{DIR_NORTH, DIR_SOUTH}},
	PT_NORTH_EAST:  {"PT_NORTH_EAST", Directions{DIR_NORTH, DIR_EAST}},
	PT_NORTH_WEST:  {"PT_NORTH_WEST", Directions{DIR_NORTH, DIR_WEST}},
	PT_SOUTH_EAST:  {"PT_SOUTH_EAST", Directions{DIR_SOUTH, DIR_EAST}},
	PT_SOUTH_WEST:  {"PT_SOUTH_WEST", Directions{DIR_SOUTH, DIR_WEST}},
	PT_EAST_WEST:   {"PT_EAST_WEST", Directions{DIR_EAST, DIR_WEST}},
	PT_START:       {"PT_START", Directions{DIR_NORTH, DIR_SOUTH, DIR_EAST, DIR_WEST}},
	PT_NONE:        {"PT_NONE", Directions{}},
}

/*
**	METHODS
 */

func (pt PipeType) CanConnect(d Direction) bool {
	return slices.Contains(allPipeTypes[pt].directions, d)
}

func (pt PipeType) ToDirections() (da Directions) {
	if ptd, ok := allPipeTypes[pt]; ok {
		return ptd.directions
	} else {
		return Directions{}
	}
}

func (pt PipeType) Name() (ptName string) {
	if ptd, ok := allPipeTypes[pt]; ok {
		return ptd.name
	} else {
		return string(pt)
	}
}

/*
**	INTERFACES
 */

// (PipeType).String() string
func (pt PipeType) String() string {
	if ptd, ok := allPipeTypes[pt]; ok {
		return ptd.name[len("PT_"):]
	} else {
		return string(pt)
	}
}

/*
**	HELPERS
 */

// StringToPipeType(pt) PipeType
func GetPipeType(ptName string) PipeType {
	for _, pt := range maps.Keys(allPipeTypes) {
		if ptName == string(pt) || ptName == pt.String() || ptName == pt.Name() {
			return pt
		}
	}
	return PT_NONE
}
