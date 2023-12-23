module github.com/nqzyx/advent-of-code/2023/day08

go 1.21.3

replace (
	github.com/nqzyx/advent-of-code/day08/nodemap => ./nodemap
	github.com/nqzyx/advent-of-code/utils => ../../internal/utils
)

require (
	github.com/nqzyx/advent-of-code/day08/nodemap v0.0.0-00010101000000-000000000000
	github.com/nqzyx/advent-of-code/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/TheAlgorithms/Go v0.0.3-alpha // indirect
	golang.org/x/exp v0.0.0-20231214170342-aacd6d4b4611 // indirect
)
