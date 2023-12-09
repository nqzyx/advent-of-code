package xref

type Entry struct {
	sources 		*Range
	destinations 	*Range
}

func NewEntry(s uint64, d uint64, l uint64) (x *Entry, err error) {
	var sources, destinations *Range
	if sources, err = NewRange(s, l); err != nil {
		return
	}
	if destinations, err = NewRange(d, l); err != nil {
		return
	}
	x = &Entry{
		sources: sources,
		destinations: destinations,
	}
	return
}

func (xe *Entry) CoversSource(v uint64) bool {
	return xe.sources.Covers(v)
}

func (xe *Entry) CoversDestination(v uint64) bool {
	return xe.destinations.Covers(v)
}


func (xe *Entry) LookupDestination(src uint64) (result uint64, ok bool) {
	var o uint64
	if o, ok = xe.sources.Offset(src); ok {
		result = xe.destinations.start + o
	}
	return
}

func (xe *Entry) LookupSource(dest uint64) (result uint64, ok bool) {
	var o uint64
	if o, ok = xe.destinations.Offset(dest); ok {
		result = xe.sources.start + o
	}
	return
}
