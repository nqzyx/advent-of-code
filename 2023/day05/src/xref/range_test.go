package xref

import (
	"reflect"
	"testing"
)

func TestNewRange(t *testing.T) {
	type In struct {
		start  uint64
		length uint64
	}
	type Want struct {
		r   *Range
		err bool
	}
	cases := []struct {
		in   In
		want Want
	}{
		{in: In{start: 2, length: 5}, want: Want{r: &Range{start: 2, length: 5, end: 5 + 2}, err: false}},
		{in: In{start: 99, length: 7}, want: Want{r: &Range{start: 99, length: 7, end: 99 + 7}, err: false}},
		{in: In{start: 0, length: 0}, want: Want{r: nil, err: true}},
	}
	for _, c := range cases {
		got, err := NewRange(c.in.start, c.in.length)
		if !reflect.DeepEqual(got, c.want.r) || c.want.err == (err == nil) {
			t.Errorf("NewRange(%v) == (%v, %v), want (%v, %v)", c.in, got, err == nil, c.want.r, c.want.err)
		}
	}
}

func TestRangeCovers(t *testing.T) {
	r := Range{start: 2, length: 5, end: 5 + 2}
	cases := []struct {
		in   uint64
		want bool
	}{
		{in: 0, want: bool(false)},
		{in: 1, want: bool(false)},
		{in: 2, want: bool(true)},
		{in: 3, want: bool(true)},
		{in: 4, want: bool(true)},
		{in: 6, want: bool(true)},
		{in: 5, want: bool(true)},
		{in: 7, want: bool(false)},
		{in: 8, want: bool(false)},
	}
	for _, c := range cases {
		got := r.Covers(c.in)
		if got != c.want {
			t.Errorf("(%v).Covers(%v) == %v, want %v", r, c.in, got, c.want)
		}

	}
}

func TestRangeOffset(t *testing.T) {
	r := Range{start: 2, length: 5, end: 5 + 2}
	type Want struct {
		n  uint64
		ok bool
	}
	cases := []struct {
		in   uint64
		want Want
	}{
		{in: 0, want: Want{n: 0, ok: false}},
		{in: 1, want: Want{n: 0, ok: false}},
		{in: 2, want: Want{n: 0, ok: true}},
		{in: 3, want: Want{n: 1, ok: true}},
		{in: 4, want: Want{n: 2, ok: true}},
		{in: 5, want: Want{n: 3, ok: true}},
		{in: 6, want: Want{n: 4, ok: true}},
		{in: 7, want: Want{n: 0, ok: false}},
		{in: 8, want: Want{n: 0, ok: false}},
	}
	for _, c := range cases {
		got, ok := r.Offset(c.in)
		if got != c.want.n || ok != c.want.ok {
			t.Errorf("(%v).Offset(%v) == %v, want %v", r, c.in, got, c.want)
		}
	}
}
