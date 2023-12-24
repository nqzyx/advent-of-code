package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
	seeds, almanac := getInput()
	answers := map[string]int64{
		"Part 1": partOne(seeds, almanac),
		"Part 2": partTwo(seeds, almanac),
	}

	if ba, err := json.MarshalIndent(answers, "", "  "); err != nil {
		panic(err)
	} else {
		fmt.Println(string(ba))
	}
}

func partOne(seeds []int64, almanac [][][]int64) int64 {
	answer := int64(math.MaxInt64)
	for _, s := range seeds {
		l := getSeedLocation(&almanac, s)
		answer = min(answer, l)
	}
	return answer
}

func partTwo(seeds []int64, almanac [][][]int64) int64 {
	seedRanges := groupNumbers(seeds, 2)
	var location int64
	for location = 0; location < 1_000_000_000; location++ {
		s := getSeedGivenLocation(almanac, location)
		if doWeHaveThatSeed(seedRanges, s) {
			return location
		}
	}
	panic(fmt.Errorf("Unable to find the closest location"))
}

func doWeHaveThatSeed(ranges [][]int64, seed int64) bool {
	for _, r := range ranges {
		start, length := r[0], r[1]
		if start <= seed && start+length > seed {
			return true
		}
	}
	return false
}

func dumpArray[T any](n string, a []T) {
	for x, i := range a {
		fmt.Printf("%v[%v] = %v\n", n, x, i)
	}
}

func getInput() (seeds []int64, almanac [][][]int64) {
	var content string // the entire file
	if ba, err := io.ReadAll(os.Stdin); err != nil {
		panic(err)
	} else {
		content = string(ba)
	}
	cc := strings.Split(content, "\n\n")
	// dumpArray("cc", cc)
	var ea [][]int64 = make([][]int64, len(cc))
	for i, c := range cc {
		ea[i] = utils.NewIntArrayFromString[int64](strings.Split(c, ":")[1])
	}
	// dumpArray("ea", ea)
	seeds = ea[0]
	almanac = make([][][]int64, len(ea[1:]))
	for i, e := range ea[1:] {
		almanac[i] = groupNumbers(e, 3)
	}
	return
}

func getSeedGivenLocation(almanac [][][]int64, step int64) int64 {
	for _, entry := range utils.Reverse(almanac) {
		for _, r := range entry {
			destination, source, length := r[0], r[1], r[2]
			if destination <= step && destination+length > step {
				step = step + (source - destination)
				break
			}
		}
	}
	return step
}

func getSeedLocation(almanac *[][][]int64, step int64) int64 {
	for _, entry := range *almanac {
		for _, rng := range entry {
			destination, source, length := rng[0], rng[1], rng[2]
			if source <= step && source+length > step {
				step = step + (source - destination)
				break
			}
		}
	}
	return step
}

func groupNumbers(a []int64, s int) [][]int64 {
	xa := make([][]int64, len(a)/s)
	i := 0
	for j := 0; j < len(a)/s; j, i = j+1, i+s {
		xa[j] = a[i : i+s]
	}
	return xa
}
