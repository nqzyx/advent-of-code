package main

import (
	"fmt"
)

// func GetBaseSeparatorCount(numGroups int) (result []int) {
// 	numSeparators := numGroups + 1
// 	result = make([]int, numSeparators)
// 	for i := range result {
// 		if i == 0 || i == numSeparators-1 {
// 			continue
// 		}
// 		result[i] = 1
// 	}
// 	return
// }

// func GetIdentityArrays(size int) (result [][]int) {
// 	result = make([][]int, 0)
// 	for i := 0; i < size; i++ {
// 		idArr := make([]int, size)
// 		idArr[i] = 1
// 		result = append(result, idArr)
// 	}
// 	return
// }

// func GetSeparatorCounts(paddingLength, numGroups int) (result [][]int) {
// 	sepCnt := numGroups + 1
// 	result = make([][]int, 0)
// 	for p := paddingLength; p >= 0; p-- {
// 		if p == 0 {
// 			result = append(result, GetBaseSeparatorCount(numGroups))
// 		}
// 		prevResult := GetSeparatorCounts(p-1, numGroups)
// 		for _, prev := range prevResult {
// 			for _, idArr := range GetIdentityArrays(sepCnt) {
// 				na := make([]int, sepCnt)
// 				for s := 0; s < sepCnt; s++ {
// 					na[s] = idArr[s] + prev[s]
// 				}
// 				result = append(result, na)
// 			}
// 		}
// 	}
// 	return
// }

// func RecursiveDots(paddingLength, numGroups int, prevResult [][]int) (result [][]int) {
// 	var dots [][]int

// 	for i := paddingLength - 1; i >= 0; i-- {
// 		if paddingLength > 1 {
// 			dots = RecursiveDots(paddingLength-1, numGroups, prevResult)
// 			for c := range dots {
// 				dots[c] = slices.Insert(dots[c], 0, i)
// 			}
// 		} else {
// 			for j := 0; j <= numGroups; j++ {
// 				start := make([]int, numGroups+1)
// 				start[j] = 1
// 				result = append(result, start)
// 			}
// 			return result
// 		}
// 	}

/*set := make(map[[]int]bool, 0, len(dots))
for _, solution := range dots {
	set[solution] = true
}
for solution := range set {
	result = append(result, solution)
}*/
// 	return
// }

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
	limit := int64(1000)
	result := make(map[int]int64, limit)
	for i := int64(0); i < limit; i++ {
		result[int(i+1)] = SumOfIntegers(i + 1)
	}
	fmt.Printf("Result: %#v\n", result)
}
