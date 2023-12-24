module github.com/nqzyx/advent-of-code/2023/day07

replace github.com/nqzyx/advent-of-code/2023/day07/camelcards => ./camelcards

replace github.com/nqzyx/advent-of-code/utils => ../../internal/utils

go 1.21.3

require (
	github.com/nqzyx/advent-of-code/2023/day07/camelcards v0.0.0-00010101000000-000000000000
	github.com/nqzyx/advent-of-code/utils v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20231214170342-aacd6d4b4611 // indirect
