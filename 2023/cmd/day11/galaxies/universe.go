package galaxies

import (
	"fmt"
	"math"

	"github.com/paulmach/orb"
)

type (
	Universe orb.MultiPoint
)

func NewUniverse(input *[]string) *Universe {
	rows, cols := len(*input), len((*input)[0])
	insertRowCount := make(map[float64]float64, rows)
	insertColCount :=  make(map[float64]float64, cols)
	u := make(Universe, 0, rows*cols/10)
	// keep track of the galaxy expansion requirements
	// assume all the rows & cols are empty
	for ir, row := range *input {
		fr := float64(ir) // convenience
		for ic, col := range row {
			fc := float64(ic)
			if col == '#' {
				g := orb.Point{fr, fc}
				// all rows from here to rows need 
				insertRowCount[ir] = true
				insertColCount[ic] = true
				// add galaxy (g) to the universe (u)
				u = append(u, g)
			}
		}
	}


	fmt.Printf("NonEmptyRows: %+v\n", insertRowCount)
	fmt.Printf("NonEmptyCols: %+v\n", insertColCount)

	fmt.Printf("Universe(orig): %v\n", u)

	for r, hasGalaxy := range insertRowCount {
		if !hasGalaxy {
			(&u).InsertRowAfter(r)
		}
	}
	// for r:=len(rowHasG)-1; r>-1; r-- {
	// 	if rowHasG[r] == true {
	// 		(&u).InsertRowAfter(r)
	// 	}
	// }
	for c, hasGalaxy := range insertColCount {
		if !hasGalaxy {
			(&u).InsertColAfter(c)
		}
	}
	fmt.Printf("Universe(exp): %v\n", u)
	uRet := Universe(u)
	return &uRet
}

func (pu *Universe) GetTotalDistance() float64 {
	u := orb.MultiPoint(*pu)
	da := make([]float64, 0, len(u)-1)
	for i:=0; i<len(*pu)-1; i++ {
		for j:=i+1; j<len(*pu); j++ {
			d := math.Abs((*pu)[i].X()-(*pu)[j].X())+math.Abs((*pu)[i].Y()-(*pu)[j].Y())
			fmt.Printf("Minimum distance from %v to %v is %v\n", (*pu)[i], (*pu)[j], d)
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
