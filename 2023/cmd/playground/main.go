package main

import (
	"fmt"

	"github.com/nqzyx/advent-of-code/utils"
)

func main() {
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

	// done := false
	// // gs := utils.MustNewGeneratorSet[uint](3, 8, signalDone)
	// gs := utils.MustNewGeneratorSet[int64](10, 16, func(gs *utils.GeneratorCollection[int64]) { done = true })
	// for {
	// 	fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)
	// 	if done {
	// 		break
	// 	}
	// }
	// fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)
	// fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)
	// fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)

	// fmt.Printf("gs: %#v\n", gs)
	done := false
	// gs := utils.MustNewGeneratorSet[uint](3, 8, signalDone)
	gs := utils.MustNewGeneratorSetSeeded[int64]([]int64{0, 1, 1, 1, 1, 1, 0}, 0, func(gs *utils.GeneratorCollection[int64]) { done = true })
	for {
		fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)
		if done {
			break
		}
	}
	fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)
	fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)
	fmt.Printf("gs.Next(): %v; done: %v\n", gs.Next(), done)

	fmt.Printf("gs: %#v\n", gs)
}
