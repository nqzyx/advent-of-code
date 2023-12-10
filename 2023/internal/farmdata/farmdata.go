package farmdata

import (
	"fmt"
	"regexp"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/utils"
	"nqzyx.xyz/advent-of-code/2023/xref"
)

type FarmData struct {
	Seeds      []uint64
	SeedRanges []xref.Range
	part1      bool
	XrefMap    map[string]xref.Xref
}

type FarmDataInterface interface {
	AddSeeds(string) *FarmData
	AddXrefMap(string) *FarmData
	Lookup(string, uint64) (uint64, error)
	Resolve(string, string, uint64) (uint64, error)
}

// Ensure Interfaces are fully implemented
var _ FarmDataInterface = &FarmData{}

func NewFarmData(stringData []string, part1 bool) *FarmData {
	FarmData := new(FarmData)
	FarmData.XrefMap = make(map[string]xref.Xref)
	FarmData.part1 = part1
	for _, data := range stringData {
		switch true {
		case strings.HasPrefix(data, "seeds:"):
			if part1 {
				FarmData.AddSeeds(data)
			} else {
				FarmData.AddSeedRanges(data)
			}
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

func (f *FarmData) AddSeedRanges(seedData string) *FarmData {
	seeds := regexp.MustCompile("^seeds: *").ReplaceAllString(seedData, "")
	seedArray := utils.NewIntArrayFromString[uint64](seeds)
	for i := 0; i < len(seedArray); i += 2 {
		rng, _ := xref.NewRange(seedArray[i], seedArray[i+1])
		f.SeedRanges = append(f.SeedRanges, *rng)
	}
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
	// fmt.Printf("xref.NewXref(%v, %v, %v, %v)\n", name, source, destination, len(xrefData))
	xref := xref.NewXref(name, source, destination, len(xrefData))
	for _, rangeData := range xrefData {
		destinationValue, sourceValue, length := func(s string) (uint64, uint64, uint64) {
			uintArray := utils.NewIntArrayFromString[uint64](s)
			return uintArray[0], uintArray[1], uintArray[2]
		}(rangeData)
		if length > 0 {
			// fmt.Printf("xref.AddRange(%v, %v, %v)\n", sourceValue, destinationValue, length)
			xref.AddRange(sourceValue, destinationValue, length)
		}
	}
	fmt.Printf("%v: %v\n", name, *xref)
	f.XrefMap[name] = *xref
	return f
}

func (f *FarmData) Lookup(name string, value uint64) (result uint64, err error) {
	if x, ok := f.XrefMap[name]; ok {
		if dv, ok := x.Lookup(x.Source, value); ok {
			return dv, nil
		}
		return 0, fmt.Errorf("\"%v\" cannot cross-reference the value %v", name, value)
	}
	return 0, fmt.Errorf("the cross-reference named \"%v\" was not found", name)
}

func (f *FarmData) Resolve(source string, destination string, value uint64) (result uint64, err error) {
	// fmt.Printf("BEGIN: Resolve(%v, %v, %v)\n", source, destination, value)
	xRefName := fmt.Sprintf("%v-to-%v", source, destination)
	if result, err = f.Lookup(xRefName, value); err == nil {
		// fmt.Printf("END: Resolve(%v, %v, %v) == %v\n", source, destination, value, result)
		return
	}
	for xrefName, x := range f.XrefMap {
		if source != x.Source {
			continue
		}
		var ok bool
		if result, ok = x.Lookup(source, value); !ok {
			return 0, fmt.Errorf(
				"failed to resolve \"%v\" to \"%v\" at \"%v\"",
				source, destination, xrefName,
			)
		}
		if destination == x.Destination {
			return
		}
		newSource, newValue := x.Destination, result
		if result, err = f.Resolve(newSource, destination, newValue); err == nil {
			// fmt.Printf("STEP: Resolve(%v, %v, %v) == %v\n", newSource, destination, newValue, result)
			return
		}
	}
	return 0, fmt.Errorf("cannot find xref from \"%v\" to \"%v\"", source, destination)
}
