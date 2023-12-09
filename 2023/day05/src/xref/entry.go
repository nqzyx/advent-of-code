package xref

type Entry struct {
	source      *Range
	destination *Range
}

func NewEntry(s uint64, d uint64, l uint64) (x *Entry, err error) {
	var source, destination *Range
	if source, err = NewRange(s, l); err != nil {
		return
	}
	if destination, err = NewRange(d, l); err != nil {
		return
	}
	x = &Entry{
		source:      source,
		destination: destination,
	}
	return
}

func (xe *Entry) CoversSource(v uint64) bool {
	return xe.source.Covers(v)
}

func (xe *Entry) CoversDestination(v uint64) bool {
	return xe.destination.Covers(v)
}

func (xe *Entry) LookupDestination(src uint64) (result uint64, ok bool) {
	var o uint64
	if o, ok = xe.source.Offset(src); ok {
		result = xe.destination.start + o
	}
	return
}

func (xe *Entry) LookupSource(dest uint64) (result uint64, ok bool) {
	var o uint64
	if o, ok = xe.destination.Offset(dest); ok {
		result = xe.source.start + o
	}
	return
}
