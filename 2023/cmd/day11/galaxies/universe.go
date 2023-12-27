package galaxies

import (
	"math"
	"strings"

	"github.com/paulmach/orb"
)

type (
	Universe orb.MultiPoint
)

func NewUniverse(input *[]string) *Universe {
	rows, cols := len(*input), len((*input)[0])
	insertRow := make(map[int]bool, rows)
	colHasGalaxy := make(map[int]bool, cols)

	u := make(Universe, 0, rows*cols/10)
	// keep track of the galaxy expansion requirements
	// assume all the rows & cols are empty
	for ir, row := range *input {
		fr := float64(ir) // convenience
		if len(strings.ReplaceAll(row, ".", "")) == 0 {
			insertRow[ir] = true
		}
		for ic, col := range row {
			fc := float64(ic)
			if col == '#' {
				colHasGalaxy[ic] = true
				g := orb.Point{fr, fc}
				u = append(u, g)
			}
		}
	}

	for r:=rows; r>=0; r-- {
		if insertRow[r] {
			u.InsertRowAfter(r)
		}
	}
	for c:=cols; c>=0; c-- {
		if hasGalaxy, ok := colHasGalaxy[c]; !ok || !hasGalaxy {
			u.InsertColAfter(c)
		}

	}

	uRet := Universe(u)
	return &uRet
}

func (pu *Universe) GetTotalDistance() float64 {
	u := orb.MultiPoint(*pu)
	da := make([]float64, 0, len(u)-1)
	for i:=0; i<len(*pu)-1; i++ {
		for j:=i+1; j<len(*pu); j++ {
			d := math.Abs((*pu)[i].X()-(*pu)[j].X())+math.Abs((*pu)[i].Y()-(*pu)[j].Y())
			da = append(da, d)
		}
	}
	var totalDistance float64
	for _, d := range da {
		totalDistance += d
	}
	return totalDistance
}

func (pu *Universe) InsertRowAfter(r int) {
	fr := float64(r)
	for i:=0; i<len(*pu); i++ {
		if (*pu)[i].X() > fr {
			(*pu)[i][0] += 1.0
		}
	}
}

func (pu *Universe) InsertColAfter(c int) {
	fc := float64(c)
	for i:=0; i<len(*pu); i++ {
		if (*pu)[i].Y() > fc {
			(*pu)[i][1] += 1.0
		}
	}	
}
