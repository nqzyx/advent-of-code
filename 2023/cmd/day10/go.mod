module github.com/nqzyx/advent-of-code/2023/day10

replace github.com/nqzyx/advent-of-code/utils => ../../internal/utils

replace github.com/nqzyx/advent-of-code/2023/day10/pipes => ./pipes

go 1.21.5

require (
	github.com/nqzyx/advent-of-code/2023/day10/pipes v0.0.0-00010101000000-000000000000
	github.com/nqzyx/advent-of-code/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/paulmach/orb v0.10.0 // indirect
	golang.org/x/exp v0.0.0-20231214170342-aacd6d4b4611 // indirect
)
