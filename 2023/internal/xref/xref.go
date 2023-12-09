package xref

import (
	"fmt"
)

type Xref struct {
	Source      string
	Destination string
	Entries     []Entry
}

type Interface interface {
	Lookup(string, string, uint64) (uint64, bool)
	AddRange(uint64, uint64, uint64) (*Xref, error)
}

var _ Interface = &Xref{}

func NewXref(name string, source string, destination string, length int) *Xref {
	xref := Xref{
		Source:      source,
		Destination: destination,
		Entries:     make([]Entry, length),
	}
	return &xref
}

func (x *Xref) Lookup(source string, destination string, value uint64) (result uint64, ok bool) {
	if source == x.Source && destination == x.Destination {
		for _, entry := range x.Entries {
			if result, ok = entry.LookupSource(value); ok {
				fmt.Printf("Lookup: (%v: %v) -> (%v: %v)\n", x.Source, value, x.Destination, result)
				return
			}
		}
		return value, false
	} else if source == x.Destination && destination == x.Source {
		for _, entry := range x.Entries {
			if result, ok = entry.LookupDestination(value); ok {
				return
			}
		}
		return value, false
	}
	return 0, false
}

func (x *Xref) AddRange(startingSource uint64, startingDestination uint64, length uint64) (*Xref, error) {
	if entry, err := NewEntry(startingSource, startingDestination, length); err != nil {
		return nil, err
	} else {
		x.Entries = append(x.Entries, *entry)
	}
	return x, nil
}
