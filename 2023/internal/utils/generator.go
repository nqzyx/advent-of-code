package utils

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type (
	TGenerator                          constraints.Integer
	GeneratorOverflowFunc[T TGenerator] func(g *Generator[T])
	GeneratorCounters                   struct {
		Value    int64
		Next     int64
		Overflow int64
	}
)

type Generator[T TGenerator] struct {
	start          T
	value          T
	maximum        T
	increment      T
	onOverflowFunc *GeneratorOverflowFunc[T]
	overflowTo     *Generator[T]
	counters       GeneratorCounters
}

const (
	// GS_SIZE_EXCEEDS_MAXIMUM   string = "the value of size (%v) must be less than the value of maximum (%v)"
	GEN_START_EXCEEDS_MAXIMUM string = "the value of start (%v) must be less than the value of maximum (%v)"
	GS_SIZE_TO_SMALL          string = "the size of the set (%v) must be greater than zero (0)"
)

func intType[T TGenerator](v T) string {
	return fmt.Sprintf("%T", v)
}

func MaxValue[T TGenerator]() T {
	var x T
	xType := intType(x)
	switch xType {
	case "uint":
		m := uint(math.MaxUint)
		return T(m)
	case "uint64":
		m := uint64(math.MaxUint64)
		return T(m)
	case "uint32":
		m := uint32(math.MaxUint32)
		return T(m)
	case "uint16":
		m := uint16(math.MaxUint16)
		return T(m)
	case "uint8":
		m := uint8(math.MaxUint8)
		return T(m)
	default:
		panic(fmt.Errorf("unsupported or unknown type: (%v)", xType))
	}
}

func NewGeneratorFull[T TGenerator](start, maximum, increment T, onOverflow *GeneratorOverflowFunc[T], overflowTo *Generator[T]) (*Generator[T], error) {
	if start > maximum {
		return nil, fmt.Errorf(GEN_START_EXCEEDS_MAXIMUM, start, maximum)
	}
	return &Generator[T]{
		start:          start,
		value:          start,
		maximum:        maximum,
		increment:      increment,
		onOverflowFunc: onOverflow,
		overflowTo:     overflowTo,
	}, nil
}

func MustNewGeneratorFull[T TGenerator](start, maximum, increment T, onOverflow *GeneratorOverflowFunc[T], overflowTo *Generator[T]) *Generator[T] {
	if g, err := NewGeneratorFull(start, maximum, increment, onOverflow, overflowTo); err != nil {
		panic(err)
	} else {
		return g
	}
}

func NewGenerator[T TGenerator](start T) (*Generator[T], error) {
	if g, err := NewGeneratorFull(start, MaxValue[T](), 1, nil, nil); err != nil {
		return nil, err
	} else {
		return g, nil
	}
}

func MustNewGenerator[T TGenerator](start T) *Generator[T] {
	if g, err := NewGenerator(start); err != nil {
		panic(err)
	} else {
		return g
	}
}

func NewGeneratorMax[T TGenerator](start, maximum T) (*Generator[T], error) {
	if g, err := NewGeneratorFull(start, maximum, 1, nil, nil); err != nil {
		return nil, err
	} else {
		return g, nil
	}
}

func MustNewGeneratorMax[T TGenerator](start, maximum T) *Generator[T] {
	if g, err := NewGeneratorMax(start, maximum); err != nil {
		panic(err)
	} else {
		return g
	}
}

func NewGeneratorMaxIncr[T TGenerator](start, maximum, increment T) (*Generator[T], error) {
	if g, err := NewGeneratorFull(start, maximum, increment, nil, nil); err != nil {
		return nil, err
	} else {
		return g, nil
	}
}

func MustNewGeneratorMaxIncr[T TGenerator](start, maximum, increment T) *Generator[T] {
	if g, err := NewGeneratorMaxIncr(start, maximum, increment); err != nil {
		panic(err)
	} else {
		return g
	}
}

func NewGeneratorOverflow[T TGenerator](start, maximum, increment T, overflowTo *Generator[T]) (*Generator[T], error) {
	if g, err := NewGeneratorFull(start, maximum, increment, nil, overflowTo); err != nil {
		return nil, err
	} else {
		return g, nil
	}
}

// This is a comment on MustNewGeneratorOverflow[T]
func MustNewGeneratorOverflow[T TGenerator](start, maximum, increment T, overflowTo *Generator[T]) *Generator[T] {
	if g, err := NewGeneratorOverflow(start, maximum, increment, overflowTo); err != nil {
		panic(err)
	} else {
		return g
	}
}

func NewGeneratorOverflowFunc[T TGenerator](start, maximum, increment T, onOverflow *GeneratorOverflowFunc[T]) (*Generator[T], error) {
	if g, err := NewGeneratorFull(start, maximum, increment, onOverflow, nil); err != nil {
		return nil, err
	} else {
		return g, nil
	}
}

func MustNewGeneratorOverflowFunc[T TGenerator](start, maximum, increment T, onOverflow *GeneratorOverflowFunc[T]) *Generator[T] {
	if g, err := NewGeneratorOverflowFunc(start, maximum, increment, onOverflow); err != nil {
		panic(err)
	} else {
		return g
	}
}

func (g *Generator[T]) Increment() T {
	return g.increment
}

func (g *Generator[T]) Max() T {
	return g.maximum
}

func (g *Generator[T]) Next() T {
	g.counters.Next++
	value := g.value
	g.value += g.increment
	if g.value > g.maximum {
		g.counters.Overflow++
		if g.onOverflowFunc != nil {
			(*g.onOverflowFunc)(g)
		}
	}
	return value
}

func (g *Generator[T]) Start() T {
	return g.start
}

func (g *Generator[T]) Value() T {
	g.counters.Value++
	return g.value
}

func (g *Generator[T]) Counters() GeneratorCounters {
	return g.counters
}
