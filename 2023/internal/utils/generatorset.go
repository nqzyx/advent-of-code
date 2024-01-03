package utils

import (
	"fmt"
)

type (
	GeneratorCollection[T TGenerator]      []*Generator[T]
	GeneratorSetOverflowFunc[T TGenerator] func(gs *GeneratorCollection[T])
	GeneratorSetCounters                   struct {
		GeneratorCounters []GeneratorCounters
		Next              int64
		Value             int64
		Overflow          int64
	}
)

type GeneratorSet[T TGenerator] struct {
	generators    GeneratorCollection[T]
	maximum       T
	lastGenerator *Generator[T]
	onOverflow    *GeneratorSetOverflowFunc[T]
	counters      GeneratorCounters
	done          bool
}

func NewGeneratorSet[T TGenerator](size T, maximum T, onOverflow GeneratorSetOverflowFunc[T]) (*GeneratorSet[T], error) {
	if size < 1 {
		return nil, fmt.Errorf(GS_SIZE_TOO_SMALL, size)
	}

	gs := new(GeneratorSet[T])
	gs.maximum = maximum
	gs.onOverflow = &onOverflow
	gs.generators = make(GeneratorCollection[T], size)

	var overflowGenerator *Generator[T]
	for i := T(0); i < size; i++ {
		if g, err := NewGeneratorFull(1, maximum-size+1, 1, gs.CreateGeneratorOverflowFunc(i), overflowGenerator); err != nil {
			return nil, err
		} else {
			gs.generators[i] = g
			overflowGenerator = g
		}
	}
	gs.lastGenerator = gs.generators[len(gs.generators)-1]
	// fmt.Printf("gs: %#v\n", gs)
	return gs, nil
}

func NewGeneratorSetSeeded[T TGenerator](seed []T, maximum T, onOverflow GeneratorSetOverflowFunc[T]) (*GeneratorSet[T], error) {
	size := T(len(seed))
	if size < 1 {
		return nil, fmt.Errorf(GS_SIZE_TOO_SMALL, size)
	}

	gs := new(GeneratorSet[T])
	gs.maximum = maximum
	gs.onOverflow = &onOverflow
	gs.generators = make(GeneratorCollection[T], size)

	var overflowGenerator *Generator[T]
	for i, s := range seed {
		if g, err := NewGeneratorFull(T(s), maximum-size+1, 1, gs.CreateGeneratorOverflowFunc(T(i)), overflowGenerator); err != nil {
			return nil, err
		} else {
			gs.generators[i] = g
			overflowGenerator = g
		}
	}
	gs.lastGenerator = gs.generators[len(gs.generators)-1]
	return gs, nil
}

func MustNewGeneratorSet[T TGenerator](size T, maximum T, onOverflow GeneratorSetOverflowFunc[T]) *GeneratorSet[T] {
	if g, err := NewGeneratorSet(size, maximum, onOverflow); err != nil {
		panic(err)
	} else {
		return g
	}
}

func MustNewGeneratorSetSeeded[T TGenerator](seed []T, maximum T, onOverflow GeneratorSetOverflowFunc[T]) *GeneratorSet[T] {
	if g, err := NewGeneratorSetSeeded(seed, maximum, onOverflow); err != nil{
		panic(err)
	} else {
		return g
	}
}

func (gs *GeneratorSet[T]) CreateGeneratorOverflowFunc(i T) *GeneratorOverflowFunc[T] {
	var f GeneratorOverflowFunc[T] = func(g *Generator[T]) {
		if gs.done {
			return
		}
		gs.counters.Overflow++
		if i == 0 {
			(*gs.onOverflow)(&gs.generators)
			gs.done = true
			return
		}
		gs.generators[i-1].Next()
		nextMaximum := gs.maximum
		for ig, g := range gs.generators {
			if T(ig) < i {
				nextMaximum -= g.value
				continue
			} else {
				g.value = g.start
				g.maximum = nextMaximum - T(len(gs.generators)-1-ig)
			}
		}
	}
	return &f
}

func (gs *GeneratorSet[T]) Max() T {
	return gs.maximum
}

func (gs *GeneratorSet[T]) Next() []T {
	if gs.done {
		return nil
	}
	gs.counters.Next++
	value := gs.Value()
	gs.lastGenerator.Next()
	return value
}

func (gs *GeneratorSet[T]) Value() []T {
	gs.counters.Value++
	value := make([]T, len(gs.generators))
	for i, g := range gs.generators {
		value[i] = g.Value()
	}
	return value
}

func (gs *GeneratorSet[T]) Counters() GeneratorCounters {
	return gs.counters
}
