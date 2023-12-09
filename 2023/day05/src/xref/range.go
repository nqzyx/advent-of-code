package xref

import "fmt"

type Range struct {
	start  uint64
	end    uint64
	length uint64
}

func NewRange(s uint64, l uint64) (x *Range, err error) {
	if l == 0 {
		return nil, fmt.Errorf("length must be greater than 0")
	}
	x = &Range{
		start:  s,
		length: l,
		end:    s + l,
	}
	return
}

func (xr *Range) Covers(v uint64) bool {
	return v >= xr.start && v < xr.end
}

func (xr *Range) Offset(v uint64) (offset uint64, ok bool) {
	if ok = xr.Covers(v); ok {
		offset = v - xr.start
	}
	return
}
