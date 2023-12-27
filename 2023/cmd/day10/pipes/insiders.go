package pipes

import (
	"fmt"
	"strings"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
	"github.com/paulmach/orb/simplify"
)

type Insiders orb.MultiPoint

func NewInsiders(p Path, tl TileList) (*Insiders, error) {
	// Do we have any data to work with?
	if len(p) == 0 {
		return nil, fmt.Errorf("the path must not be empty")
	}
	if len(tl) == 0 {
		return nil, fmt.Errorf("the tile list must not be empty")
	}
	if !p.Closed() {
		return nil, fmt.Errorf("the path must return to the starting point")
	}
	epsilon := 0.25
	// Do we have any candidates for insiders?
	// Reduce the path to it's vertices (aka, corners or turns)
	pSimple := simplify.DouglasPeucker(epsilon).Simplify(orb.Ring(p).Clone()).(orb.Ring)
	tInsiders := make(orb.MultiPoint, 0, len(tl))
	// check the tiles for candidates
	for _, t := range tl {
		tP := orb.Point(t.Location)
		if !t.OnPath && // tiles on the Path are not are NOT candidates
			planar.RingContains(pSimple, tP) {
			tInsiders = append(tInsiders, orb.Point(t.Location))
		}
	}
	if len(tInsiders) == 0 {
		return &Insiders{}, nil
	}
	return (*Insiders)(&tInsiders), nil
}

func (i Insiders) String() string {
	mp := orb.MultiPoint(i)
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("{Length:%v,[", len(mp)))
	for _, p := range mp {
		sb.WriteString(fmt.Sprintf("{R:%v,C:%v}", p.X(), p.Y()))
	}
	sb.WriteString("]}")
	return sb.String()
}
