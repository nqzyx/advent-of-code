module github.com/nqzyx/advent-of-code/2023/day09

go 1.21.3

replace github.com/nqzyx/advent-of-code/utils => ../../internal/utils

replace github.com/nqzyx/advent-of-code/2023/day09/oasis => ./oasis

require (
	github.com/nqzyx/advent-of-code/day09/oasis v0.0.0-00010101000000-000000000000
	github.com/nqzyx/advent-of-code/utils v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20231214170342-aacd6d4b4611 // indirect
