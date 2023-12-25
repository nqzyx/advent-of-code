package pipes

import "github.com/paulmach/orb"

type Path     struct {
	Steps orb.Ring
	Sides orb.MultiLineString
}

func (p Path) AddTile(t *Tile) {
	p.AddLocation(t.Location)
}

func (p *Path) AddLocation(l Location) {
	p.Steps = append(p.Steps, orb.Point(l))
}