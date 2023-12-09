package xref

import (
	"reflect"
	"testing"
)

func TestNewRange(t *testing.T) {
	type In struct {
		Start  uint64
		Length uint64
	}
	type Want struct {
		R      *Range
		HasErr bool
	}
	cases := []struct {
		In   In
		Want Want
	}{
		{In: In{Start: 2, Length: 5}, Want: Want{R: &Range{Start: 2, End: 2 + 5}, HasErr: false}},
		{In: In{Start: 99, Length: 7}, Want: Want{R: &Range{Start: 99, End: 99 + 7}, HasErr: false}},
		{In: In{Start: 0, Length: 0}, Want: Want{R: nil, HasErr: true}},
	}
	for _, c := range cases {
		in, want := c.In, c.Want
		got, err := NewRange(in.Start, in.Length)
		if !reflect.DeepEqual(got, want.R) || want.HasErr == (err == nil) {
			t.Errorf("NewRange(%v) == (%v, %v), want (%v, %v)", in, got, err == nil, want.R, want.HasErr)
		}
	}
}

func TestRangeCovers(t *testing.T) {
	r := Range{Start: 2, End: 5 + 2}
	cases := []struct {
		In   uint64
		Want bool
	}{
		{In: 0, Want: bool(false)},
		{In: 1, Want: bool(false)},
		{In: 2, Want: bool(true)},
		{In: 3, Want: bool(true)},
		{In: 4, Want: bool(true)},
		{In: 6, Want: bool(true)},
		{In: 5, Want: bool(true)},
		{In: 7, Want: bool(false)},
		{In: 8, Want: bool(false)},
	}
	for _, c := range cases {
		in, want := c.In, c.Want
		got := r.Covers(in)
		if got != want {
			t.Errorf("(%v).Covers(%v) == %v, want %v", r, in, got, want)
		}

	}
}

func TestRangePosition(t *testing.T) {
	r := Range{Start: 2, End: 5 + 2}
	type Want struct {
		N  uint64
		OK bool
	}
	cases := []struct {
		In   uint64
		Want Want
	}{
		{In: 0, Want: Want{N: 0, OK: false}},
		{In: 1, Want: Want{N: 0, OK: false}},
		{In: 2, Want: Want{N: 0, OK: true}},
		{In: 3, Want: Want{N: 1, OK: true}},
		{In: 4, Want: Want{N: 2, OK: true}},
		{In: 5, Want: Want{N: 3, OK: true}},
		{In: 6, Want: Want{N: 4, OK: true}},
		{In: 7, Want: Want{N: 0, OK: false}},
		{In: 8, Want: Want{N: 0, OK: false}},
	}
	for _, c := range cases {
		in, want := c.In, c.Want
		got, ok := r.Position(in)
		if got != want.N || ok != want.OK {
			t.Errorf("(%v).Position(%v) == %v, want %v", r, in, got, want)
		}
	}
}
