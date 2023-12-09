package xref

type Xref struct {
	source      string
	destination string
	entries     []Entry
}

// Ensure Interfaces are implemented
var _ Interface = &Xref{}

func NewXref(name string, source string, destination string, size int) *Xref {
	xref := Xref{
		source:      source,
		destination: destination,
		entries:     make([]Entry, 0),
	}
	return &xref
}

func (x *Xref) Source() string {
	return x.source
}
func (x *Xref) SetSource(source string) *Xref {
	x.source = source
	return x
}

func (x *Xref) Destination() string {
	return x.destination
}
func (x *Xref) SetDestination(destination string) *Xref {
	x.destination = destination
	return x
}

func (x *Xref) Lookup(s string, v uint64) (result uint64, ok bool) {
	switch s {
	case x.source:
		for _, r := range x.entries {
			if result, ok = r.LookupSource(v); ok {
				return
			}
		}
		return v, true
	case x.destination:
		for _, r := range x.entries {
			if result, ok = r.LookupDestination(v); ok {
				return
			}
		}
		return v, true
	default:
		return 0, false
	}
}

func (x *Xref) AddRange(s1 uint64, d1 uint64, l uint64) (*Xref, error) {
	if e, err := NewEntry(s1, d1, l); err != nil {
		return x, err
	} else {
		x.entries = append(x.entries, *e)
	}
	return x, nil
}
