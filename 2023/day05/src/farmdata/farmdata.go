package farmdata

import (
	"fmt"
	"regexp"
	"strings"

	"nqzyx.xyz/advent-of-code/2023/day05/utils"
	"nqzyx.xyz/advent-of-code/2023/day05/xref"
)

type FarmData struct {
	seeds   []uint64
	xRefMap map[string]xref.Xref
}

type FarmDataInterface interface {
	Seeds() []uint64
	AddSeeds(string) *FarmData
	AddXrefMap(string) *FarmData
	DestinationValue(string, uint64) (uint64, error)
	DestinationValueByType(string, uint64, string) (uint64, error)
}

// Ensure Interfaces are fully implemented
var _ FarmDataInterface = &FarmData{}

func NewFarmData(sArr []string) *FarmData {
	FarmData := new(FarmData)
	FarmData.xRefMap = make(map[string]xref.Xref)

	for _, data := range sArr {
		switch true {
		case strings.HasPrefix(data, "seeds:"):
			FarmData.AddSeeds(data)
		default:
			FarmData.AddXrefMap(data)
		}
	}
	return FarmData
}

func (fd *FarmData) Seeds() []uint64 {
	return fd.seeds
}

func (fd *FarmData) AddSeeds(s string) *FarmData {
	seeds := regexp.MustCompile("^seeds: *").ReplaceAllString(s, "")
	fd.seeds = utils.NewIntArrayFromString[uint64](seeds)
	return fd
}

func (fd *FarmData) AddXrefMap(xrefData string) *FarmData {
	name, data := func()(string, []string){d := strings.Split(xrefData, "\n"); return d[0], d[1:]}() 
	name = strings.TrimSpace(strings.ReplaceAll(name, "map:", ""))
	source, destination := func()(string, string){ t := strings.Split(name, "-to-"); return t[0], t[1]}()
	xref := xref.NewXref(name, source, destination, len(data))
	for _, s := range data {
		ia := utils.NewIntArrayFromString[uint64](s)
		xref.AddRange(ia[0], ia[1], ia[2])
	}
	fd.xRefMap[name] = *xref
	return fd
}

func (fd *FarmData) DestinationValue(s string, v uint64) (result uint64, err error) {
	if x, ok := fd.xRefMap[s]; ok {
		if dv, ok := x.Lookup(x.Source(), v); ok {
			return dv, nil
		}
		return 0, fmt.Errorf("cannot find value for \"%v\" = %v", s, v)
	}
	return 0, fmt.Errorf("xref name \"%v\" was not found", s)
}

func (fd *FarmData) DestinationValueByType(s string, v uint64, d string) (result uint64, err error) {
	xRefName := fmt.Sprintf("%v-to-%v", s, d)
	if result, err = fd.DestinationValue(xRefName, v); err == nil {
		return
	}
	var x xref.Xref
	for _, x = range fd.xRefMap {
		if x.Source() == s {
			if newValue, ok := x.Lookup(s, v); ok {
				result, err = fd.DestinationValueByType(
					x.Destination(), 
					newValue,
					d,
				)
				return
			}
		}
	}
	return 0, fmt.Errorf("cannot find xref from \"%v\" to \"%v\"", s, d)
}
