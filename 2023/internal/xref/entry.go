package xref

type Entry struct {
	Source      Range
	Destination Range
}

type EntryInterface interface {
	LookupDestination(uint64) (uint64, bool)
	LookupSource(uint64) (uint64, bool)
}

var _ EntryInterface = &Entry{}

func NewEntry(startSource uint64, startDestination uint64, length uint64) (e *Entry, err error) {
	var source, destination *Range
	if source, err = NewRange(startSource, length); err != nil {
		return
	}
	if destination, err = NewRange(startDestination, length); err != nil {
		return
	}
	e = &Entry{
		Source:      *source,
		Destination: *destination,
	}
	return
}

func (e *Entry) LookupDestination(sourceValue uint64) (result uint64, ok bool) {
	var position uint64
	if position, ok = e.Source.Position(sourceValue); ok {
		result = e.Destination.Start + position
		return
	}
	return 0, false
}

func (e *Entry) LookupSource(destinationValue uint64) (result uint64, ok bool) {
	var position uint64
	if position, ok = e.Destination.Position(destinationValue); ok {
		result = e.Source.Start + position
		return
	}
	return 0, false
}
