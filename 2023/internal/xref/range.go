package xref

import (
	"fmt"
)

type Range struct {
	Start uint64
	End   uint64
}

type RangeInterface interface {
	Length() uint64
	Covers(uint64) bool
	Position(uint64) (uint64, bool)
}

var _ RangeInterface = &Range{}

func NewRange(start uint64, length uint64) (r *Range, err error) {
	if length == 0 {
		return nil, fmt.Errorf("length must be greater than 0")
	}
	r = &Range{
		Start: start,
		End:   start + length,
	}
	return
}

func (r *Range) Length() uint64 {
	return r.End - r.Start
}

func (r *Range) Covers(value uint64) bool {
	return value >= r.Start && value < r.End
}

func (r *Range) Position(value uint64) (position uint64, ok bool) {
	if ok = r.Covers(value); ok {
		position = value - r.Start
	}
	return
}
