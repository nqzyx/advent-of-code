package pipes

import (
	"fmt"

	"github.com/paulmach/orb"
)

type Location orb.Point

// Methods

func (l Location) GetRowCol() (r, c int) {
	r, c = int(l[0]), int(l[1])
	return
}

// Interfaces

func (l Location) String() string {
	p := orb.Point(l)
	return fmt.Sprintf("[R:%v,C:%v]", p.X(), p.Y())
}

// Helpers

func GetLocation(r, c int) (l Location) {
	l = Location{float64(r), float64(c)}
	return
}
