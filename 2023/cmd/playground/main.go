package main

import (
	"fmt"
	"slices"

	"github.com/nqzyx/advent-of-code/utils"
)

func GetBaseSeparatorCount(numGroups int) (result []int) {
	numSeparators := numGroups + 1
	result = make([]int, numSeparators)
	for i := range result {
		if i == 0 || i == numSeparators-1 {
			continue
		}
		result[i] = 1
	}
	return
}

func GetIdentityArrays(size int) (result [][]int) {
	result = make([][]int, 0)
	for i := 0; i < size; i++ {
		idArr := make([]int, size)
		idArr[i] = 1
		result = append(result, idArr)
	}
	return
}

func GetSeparatorCounts(paddingLength, numGroups int) (result [][]int) {
	sepCnt := numGroups + 1
	result = make([][]int, 0)
	for p := paddingLength; p >= 0; p-- {
		if p == 0 {
			result = append(result, GetBaseSeparatorCount(numGroups))
		}
		prevResult := GetSeparatorCounts(p-1, numGroups)
		for _, prev := range prevResult {
			for _, idArr := range GetIdentityArrays(sepCnt) {
				na := make([]int, sepCnt)
				for s := 0; s < sepCnt; s++ {
					na[s] = idArr[s] + prev[s]
				}
				result = append(result, na)
			}
		}
	}
	return
}

func RecursiveDots(paddingLength, numGroups int, prevResult [][]int) (result [][]int) {
	var dots [][]int

	for i := paddingLength - 1; i >= 0; i-- {
		if paddingLength > 1 {
			dots = RecursiveDots(paddingLength-1, numGroups, prevResult)
			for c := range dots {
				dots[c] = slices.Insert(dots[c], 0, i)
			}
		} else {
			for j := 0; j <= numGroups; j++ {
				start := make([]int, numGroups+1)
				start[j] = 1
				result = append(result, start)
			}
			return result
		}
	}

	/*set := make(map[[]int]bool, 0, len(dots))
	  for _, solution := range dots {
	  	set[solution] = true
	  }
	  for solution := range set {
	  	result = append(result, solution)
	  }*/
	return
}

func SumOfIntegers(a int64) int64 {
	var l int64 = int64(1)
	var n int64 = a
	return n * (a + l) / int64(2)
}

func main() {
	// result := RecursiveDots(2, 3)
	// result := GetBaseSeparatorCount(7)
	// result := GetIdentityArrays(5)
	// result := GetSeparatorCounts(3, 5)
	// fmt.Printf("%v\n", result)
	// limit := int64(1_000_000_000)
	// var result int64
	// for i := int64(1_000_000); i <= limit; i += 1_000_000 {
	// 	result = SumOfIntegers(i)
	// 	fmt.Printf("SumOfIntegers(%v): %#v\n", i, result)
	// }
	// fmt.Printf("MustNewGenerator(0,  1,  10): %v\n", utils.MustNewGenerator[uint](0))
	// for i := 0; i < 11; i++ {
	// }
	// fmt.Printf("MustNewGenerator(0, 10, 100): %v\n", utils.MustNewGeneratorMaxIncr[uint](0, 10, 100))

	// gTens := utils.MustNewGeneratorMaxIncr[uint](0, 100, 10)
	// gOnes := utils.MustNewGeneratorOverflow[uint](0, 10, 1, gTens)

	// fmt.Printf("gTens: %#v\n", gTens)
	// fmt.Printf("gOnes: %#v\n", gOnes)

	// for i := 0; i < 110; i++ {
	// 	// var ones int
	// 	// if i == 0 {
	// 	// 	ones = gOnes.Value()
	// 	// } else {
	// 	//  	ones = gOnes.Next()
	// 	// }
	// 	ones := gOnes.Next()
	// 	tens := gTens.Value()
	// 	fmt.Printf("i: %v; tens: %v; ones %v\n", i, tens, ones)
	// }

	// g := utils.MustNewGenerator[uint](0)
	// fmt.Println("g: ", g)
	done := false
	// signalDone := func(g *utils.Generator[uint]) {
	// 	done = true
	// }

	// gs := utils.MustNewGeneratorSet[uint](3, 8, signalDone)
	gs := utils.MustNewGeneratorSet[uint](6, 10, func(gs *utils.GeneratorCollection[uint]) { done = true })
	for i := 0; i < 300; i++ {
		fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)
		if done {
			break
		}
	}
	fmt.Printf("gs: %#v\n", gs)
}
