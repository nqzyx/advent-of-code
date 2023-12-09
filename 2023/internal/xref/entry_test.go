package xref

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestNewEntry(t *testing.T) {
	type In struct {
		s uint64
		d uint64
		l uint64
	}
	type Want struct {
		e   *Entry
		err bool
	}
	cases := []struct {
		in   In
		want Want
	}{
		{
			in: In{s: 2, d: 99, l: 5},
			want: Want{e: &Entry{
				Source:      Range{Start: 2, End: 5 + 2},
				Destination: Range{Start: 99, End: 99 + 5},
			}, err: false},
		},
		{
			in: In{s: 99, d: 2, l: 7},
			want: Want{
				e: &Entry{
					Source:      Range{Start: 99, End: 99 + 7},
					Destination: Range{Start: 2, End: 2 + 7},
				}, err: false},
		},
		{
			in: In{s: 0, d: 0, l: 0},
			want: Want{
				e:   nil,
				err: true,
			},
		},
	}
	for _, c := range cases {
		got, err := NewEntry(c.in.s, c.in.d, c.in.l)
		if !reflect.DeepEqual(&got, &c.want.e) || (err == nil) == c.want.err {
			t.Errorf("NewEntry(%v) == (%v, %v), want (%v, %v)", c.in, *got, err == nil, *c.want.e, c.want.err)
		}
	}
}

func TestCovers(t *testing.T) {
	r := Range{Start: 2, End: 5 + 2}
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

func TestPosition(t *testing.T) {
	r := Range{Start: 2, End: 5 + 2}
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
		got, ok := r.Position(c.in)
		if got != c.want.n || ok != c.want.ok {
			t.Errorf("(%v).Position(%v) == %v, want %v", r, c.in, got, c.want)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	type localRange Range
	var r localRange = localRange{
		Start: 1234,
		End:   5678,
	}
	if json, err := json.Marshal(r); err != nil {
		panic(err)
	} else {
		fmt.Println("json:", json)
	}
}
