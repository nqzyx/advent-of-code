package main

import (
	"encoding/json"
)

// Methods
type XRefer interface {
	// Sources
	SourceName() string
	SetSourceName(sourceName string) *XRef
	GetSourceValue(target uint32) uint32
	// Targets
	TargetName() string
	SetTargetName(targetName string) *XRef
	GetTargetValue(source uint32) uint32
	// XRef Ranges
	AddXRefRange(firstSource uint32, firstTarget uint32, entries uint32) *XRef
}

// Properties
type XRef struct {
	XRefer               // interface enforcement
	sourceName           string
	targetName           string
	sourceToTargetValues XRefEntries
	targetToSourceValues XRefEntries
}

// Interface Enforcement
var _ XRefer = XRef{}
var _ json.Marshaler = XRef{}

type XRefEntries map[uint32]uint32

// func (entries XRefEntries) String() (outString string) {
// 	outString = fmt.Sprintf(
// 		"[ %v ]", func() (outString string) {
// 			for key, value := range entries {
// 				outString += fmt.Sprintf("{ %v: %v },", key, value)
// 			}
// 			outString = outString[0:len(outString)-1] + " }"
// 			return
// 		}(),
// 	)
// 	outString = "["
// 	for sourceValue, targetValue := range entries {
// 		outString += fmt.Sprintf("{%v: %v},", sourceValue, targetValue)
// 	}
// 	outString = outString[0:len(outString)-1] + "]"
// 	return
// }

// func (xref XRef) String() (outString string) {
// 	outString = fmt.Sprintf(
// 		"{ sourceName: \"%v\", targetName: \"%v\", sourceToTargetValues: %v, targetToSourceValues: %v",
// 		xref.sourceName, xref.targetName, xref.sourceToTargetValues.String(), xref.targetToSourceValues.String(),
// 	)
// 	return
// }

func NewXRef(sourceName string, targetName string) *XRef {
	xref := XRef{
		sourceName:           sourceName,
		targetName:           targetName,
		sourceToTargetValues: make(XRefEntries),
		targetToSourceValues: make(XRefEntries),
	}
	return &xref
}

// Sources
func (c XRef) SourceName() string {
	return c.sourceName
}
func (c XRef) SetSourceName(sourceName string) *XRef {
	c.sourceName = sourceName
	return &c
}
func (c XRef) GetSourceValue(target uint32) uint32 {
	if value, ok := c.targetToSourceValues[target]; ok {
		return value
	} else {
		return target
	}
}

// Targets
func (c XRef) TargetName() string {
	return c.targetName
}
func (c XRef) SetTargetName(targetName string) *XRef {
	c.targetName = targetName
	return &c
}
func (c XRef) GetTargetValue(source uint32) uint32 {
	if value, ok := c.sourceToTargetValues[source]; ok {
		return value
	} else {
		return source
	}
}

// XRefs

func (c XRef) AddXRefRange(firstSource uint32, firstTarget uint32, entries uint32) *XRef {
	for entry, source, target := uint32(0), firstSource, firstTarget; entry < entries; entry, source, target = entry+1, source+1, target+1 {
		c.sourceToTargetValues[source] = target
		c.targetToSourceValues[target] = source
	}
	return &c
}

func (c XRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}
