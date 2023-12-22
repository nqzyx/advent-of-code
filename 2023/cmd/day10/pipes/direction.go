package pipes

type Direction string

const (
	Unknown Direction = "U"
	North   Direction = "N"
	South   Direction = "S"
	East    Direction = "E"
	West    Direction = "W"
)

type Directions map[Direction]bool
