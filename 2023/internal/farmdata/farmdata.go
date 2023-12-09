package farmdata

import (
	"fmt"
	"regexp"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/utils"
	"nqzyx.xyz/advent-of-code/2023/xref"
)

type FarmData struct {
	Seeds   []uint64
	XrefMap map[string]xref.Xref
}

type FarmDataInterface interface {
	AddSeeds(string) *FarmData
	AddXrefMap(string) *FarmData
	DestinationValue(string, uint64) (uint64, error)
	DestinationValueByType(string, uint64, string) (uint64, error)
}

// Ensure Interfaces are fully implemented
var _ FarmDataInterface = &FarmData{}

func NewFarmData(stringData []string) *FarmData {
	FarmData := new(FarmData)
	FarmData.XrefMap = make(map[string]xref.Xref)

	for _, data := range stringData {
		switch true {
		case strings.HasPrefix(data, "seeds:"):
			FarmData.AddSeeds(data)
		default:
			FarmData.AddXrefMap(data)
		}
	}
	return FarmData
}

func (f *FarmData) AddSeeds(seedData string) *FarmData {
	seeds := regexp.MustCompile("^seeds: *").ReplaceAllString(seedData, "")
	f.Seeds = utils.NewIntArrayFromString[uint64](seeds)
	return f
}

func (f *FarmData) AddXrefMap(xrefMapData string) *FarmData {
	name, xrefData := func() (string, []string) {
		stringArray := strings.Split(xrefMapData, "\n")
		return strings.TrimSpace(strings.ReplaceAll(stringArray[0], "map:", "")), stringArray[1:]
	}()
	source, destination := func() (string, string) {
		stringArray := strings.Split(name, "-to-")
		return stringArray[0], stringArray[1]
	}()
	xref := xref.NewXref(name, source, destination, len(xrefData))
	for _, rangeData := range xrefData {
		destinationValue, sourceValue, length := func(s string) (uint64, uint64, uint64) {
			uintArray := utils.NewIntArrayFromString[uint64](s)
			return uintArray[0], uintArray[1], uintArray[2]
		}(rangeData)
		xref.AddRange(sourceValue, destinationValue, length)
	}
	f.XrefMap[name] = *xref
	return f
}

func (f *FarmData) DestinationValue(s string, v uint64) (result uint64, err error) {
	if x, ok := f.XrefMap[s]; ok {
		if dv, ok := x.Lookup(x.Source, x.Destination, v); ok {
			return dv, nil
		}
		return 0, fmt.Errorf("cannot find value for \"%v\" = %v", s, v)
	}
	return 0, fmt.Errorf("xref name \"%v\" was not found", s)
}

func (f *FarmData) DestinationValueByType(s string, v uint64, d string) (result uint64, err error) {
	xRefName := fmt.Sprintf("%v-to-%v", s, d)
	if result, err = f.DestinationValue(xRefName, v); err == nil {
		return
	}
	var x xref.Xref
	for _, x = range f.XrefMap {
		if x.Source == s {
			if newValue, ok := x.Lookup(s, v); ok {
				result, err = f.DestinationValueByType(
					x.Destination,
					newValue,
					d,
				)
				return
			}
		}
	}
	return 0, fmt.Errorf("cannot find xref from \"%v\" to \"%v\"", s, d)
}
