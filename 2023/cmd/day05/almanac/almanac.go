package almanac

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nqzyx/advent-of-code/utils"
	"github.com/nqzyx/advent-of-code/xref"
)

type Almanac struct {
	Seeds      []uint64
	SeedRanges []xref.Range
	part1      bool
	XrefMap    map[string]xref.Xref
}

type AlmanacInterface interface {
	AddSeeds(string) *Almanac
	AddSeedRanges(string) *Almanac
	AddXrefMap(string) *Almanac
	Lookup(string, uint64) (uint64, error)
	Resolve(string, string, uint64) (uint64, error)
}

// Ensure Interfaces are fully implemented
var _ AlmanacInterface = &Almanac{}

func NewAlmanac(stringData []string, part1 bool) *Almanac {
	Almanac := new(Almanac)
	Almanac.XrefMap = make(map[string]xref.Xref)
	Almanac.part1 = part1
	for _, data := range stringData {
		switch true {
		case strings.HasPrefix(data, "seeds:"):
			if part1 {
				Almanac.AddSeeds(data)
			} else {
				Almanac.AddSeedRanges(data)
			}
		default:
			Almanac.AddXrefMap(data)
		}
	}
	return Almanac
}

func (f *Almanac) AddSeeds(seedData string) *Almanac {
	seeds := regexp.MustCompile("^seeds: *").ReplaceAllString(seedData, "")
	f.Seeds = utils.NewNumericArrayFromString[uint64](seeds)
	return f
}

func (f *Almanac) AddSeedRanges(seedData string) *Almanac {
	seeds := regexp.MustCompile("^seeds: *").ReplaceAllString(seedData, "")
	seedArray := utils.NewNumericArrayFromString[uint64](seeds)
	for i := 0; i < len(seedArray); i += 2 {
		rng, _ := xref.NewRange(seedArray[i], seedArray[i+1])
		f.SeedRanges = append(f.SeedRanges, *rng)
	}
	return f
}

func (f *Almanac) AddXrefMap(xrefMapData string) *Almanac {
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
			uintArray := utils.NewNumericArrayFromString[uint64](s)
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

func (f *Almanac) Lookup(name string, value uint64) (result uint64, err error) {
	if x, ok := f.XrefMap[name]; ok {
		if dv, ok := x.Lookup(x.Source, value); ok {
			return dv, nil
		}
		return 0, fmt.Errorf("\"%v\" cannot cross-reference the value %v", name, value)
	}
	return 0, fmt.Errorf("the cross-reference named \"%v\" was not found", name)
}

func (f *Almanac) Resolve(source string, destination string, value uint64) (result uint64, err error) {
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
