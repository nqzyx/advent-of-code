package pipes

import (
	"fmt"
	"strings"

	"github.com/paulmach/orb"
)

type Path orb.LineString

func NewPath() *Path {
	p := new(Path)
	*p = Path(make(orb.LineString, 0))
	return p
}

func (p *Path) Add(t *Tile) {
	p.AddLocation(t.Location)
}

func (p *Path) AddLocation(l Location) {
	*p = append(*p, orb.Point(l))
}

func (p *Path) Closed() bool {
	return orb.Ring(*p).Closed()
}

func (p Path) String() string {
	mp := orb.LineString(p)
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("{Length:%v,[", len(mp)))
	for _, p := range mp {
		sb.WriteString(fmt.Sprintf("{R:%v,C:%v}", p.X(), p.Y()))
	}
	sb.WriteString("]}")
	return sb.String()
}
