package main

type XRef struct {
	sourceName           string
	targetName           string
	sourceToTargetValues XRefEntries
	targetToSourceValues XRefEntries
}

type XRefEntries map[uint32]uint32

type XRefr interface {
	// constructor
	NewXRef(sourceName string, targetName string) *XRef
	// getters
	SourceName() string
	TargetName() string
	GetTargetValue(source uint32) uint32
	GetSourceValue(target uint32) uint32

	// setters
	SetSourceName(sourceName string) *XRef
	SetTargetName(targetName string) *XRef
	AddXRefRange(firstSource uint32, firstTarget uint32, entries uint32) *XRef
}

func NewXRef(sourceName string, targetName string) *XRef {
	xref := &XRef{
		sourceName:           sourceName,
		targetName:           targetName,
		sourceToTargetValues: make(XRefEntries),
		targetToSourceValues: make(XRefEntries),
	}
	return xref
}

func (c XRef) SourceName() string {
	return c.sourceName
}
func (c XRef) TargetName() string {
	return c.targetName
}

func (c *XRef) SetSourceName(sourceName string) *XRef {
	c.sourceName = sourceName
	return c
}
func (c *XRef) SetTargetName(targetName string) *XRef {
	c.targetName = targetName
	return c
}

func (c *XRef) AddXRefRange(firstSource uint32, firstTarget uint32, entries uint32) *XRef {
	for entry, source, target := uint32(0), firstSource, firstTarget; entry < entries; entry, source, target = entry+1, source+1, target+1 {
		c.sourceToTargetValues[source], c.targetToSourceValues[target] = target, source
	}
	return c
}

func (c XRef) GetTargetValue(source uint32) uint32 {
	if value, ok := c.sourceToTargetValues[source]; ok {
		return value
	} else {
		return source
	}
}

func (c XRef) GetSourceValue(target uint32) uint32 {
	if value, ok := c.targetToSourceValues[target]; ok {
		return value
	} else {
		return target
	}
}
