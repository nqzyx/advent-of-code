package pipes

// Location helpers

func GetPoint(r, c int) (l Location) {
	l = Location{float64(r), float64(c)}
	return
}

func GetRowCol(l Location) (r, c int) {
	r, c = int(l[0]), int(l[1])
	return
}
