module github.com/nqzyx/advent-of-code/2023/day05

go 1.21.3

replace github.com/nqzyx/advent-of-code/utils => ../../internal/utils

require github.com/nqzyx/advent-of-code/2023/day05/almanac v0.0.0-00010101000000-000000000000

require github.com/nqzyx/advent-of-code/utils v0.0.0-00010101000000-000000000000

require (
	github.com/nqzyx/advent-of-code/xref v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/exp v0.0.0-20231214170342-aacd6d4b4611 // indirect
)

replace github.com/nqzyx/advent-of-code/2023/day05/almanac => ./almanac

replace github.com/nqzyx/advent-of-code/xref => ../../internal/xref
