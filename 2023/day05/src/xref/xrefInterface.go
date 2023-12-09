package xref

type Interface interface {
	Source() string
	SetSource(string) *Xref
	// Targets
	Destination() string
	SetDestination(string) *Xref
	Lookup(string, uint64) (uint64, bool)
	AddRange(uint64, uint64, uint64) (*Xref, error)
}