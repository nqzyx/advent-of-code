package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

var integerRegexp = regexp.MustCompile("[[:digit:]]")

func NewIntArrayFromString[T constraints.Integer](s string) (result []T) {
	s = strings.TrimSpace(s)
	sArr := integerRegexp.FindAllString(s, -1)
	result = make([]T, len(sArr))
	var x T
	b := int(reflect.TypeOf(x).Size()) * 8
	for i, v := range sArr {
		x, err := (strconv.ParseInt(v, 0, b))
		if err != nil {
			panic(err)
		}
		result[i] = T(x)
	}
	return
}

// func makeUInt32Array(s string, splitter string) (result []uint32) {
// 	this := strings.TrimSpace(string(s))
// 	space, space2 := " ", "  "
// 	for strings.Contains(this, space2) {
// 		this = strings.ReplaceAll(this, space2, space)
// 	}
// 	theseStrings := strings.Split(this, splitter)
// 	result = make([]uint32, len(theseStrings))
// 	for thisIndex, thisString := range theseStrings {
// 		thisValue, err := strconv.ParseInt(thisString, 10, 32)
// 		if err != nil {
// 			panic(err)
// 		}
// 		result[thisIndex] = uint32(thisValue)
// 	}
// 	return
// }

type XRefMap map[string]XRef

type GardenData struct {
	seeds   []uint32
	xRefMap XRefMap
}

func NewGardenData(inputData []string) *GardenData {
	gardenData := new(GardenData)
	gardenData.xRefMap = make(map[string]XRef)

	for _, data := range inputData {
		switch true {
		case strings.HasPrefix(data, "seeds:"):
			gardenData.SetSeeds(data)
		default:
			gardenData.AddXrefMap(data)
		}
	}
	return gardenData
}

// func (m XRefMap) String() string {
// 	return fmt.Sprintf(
// 		"[ %v ]",
// 		func() (result string) {
// 			for key, value := range m {
// 				result += fmt.Sprintf(
// 					"%v: %v,",
// 					key,
// 					value.String(),
// 				)
// 			}
// 			result = result[0 : len(result)-1]
// 			return
// 		}(),
// 	)

// }

func (gd GardenData) Seeds() []uint32 {
	return gd.seeds
}

func (gd GardenData) SetSeeds(data string) *GardenData {
	data = regexp.MustCompile("(seeds: *)").ReplaceAllString(data, "")
	gd.seeds = NewIntArrayFromString[uint32](data)
	return &gd
}

func (gd GardenData) AddXrefMap(s string) *GardenData {
	sParts := strings.Split(s, "\n")
	xrefName := strings.ReplaceAll(sParts[0], " map:", "")
	names := strings.Split(xrefName, "-to-")
	sourceName, targetName := names[0], names[1]
	xref := *NewXRef(sourceName, targetName)
	for _, data := range sParts[1:] {
		ia := NewIntArrayFromString[uint32](data)
		xref.AddXRefRange(ia[0], ia[1], ia[2])
	}
	gd.xRefMap[xrefName] = xref
	return &gd
}

func (gd GardenData) GetTargetValue(sourceName string, sourceValue uint32, targetName string) (targetValue uint32) {

	xRefName := fmt.Sprintf("%v-to-%v", sourceName, targetName)
	if xref, ok := gd.xRefMap[xRefName]; ok {
		targetValue = xref.GetTargetValue(sourceValue)
		jsonPrint(map[string]string{"sourceName": sourceName, "sourceValue": fmt.Sprint(sourceValue), "targetName": targetName, "targetValue": fmt.Sprint(targetValue)})
		return
	}
	for _, xref := range gd.xRefMap {
		if xref.SourceName() == sourceName {
			targetValue = gd.GetTargetValue(xref.TargetName(), xref.GetTargetValue(sourceValue), targetName)
			jsonPrint(map[string]string{"sourceName": sourceName, "sourceValue": fmt.Sprint(sourceValue), "targetName": targetName, "targetValue": fmt.Sprint(targetValue)})
			return
		}
	}
	panic(fmt.Sprintf("Cannot find xref from %v to %v", sourceName, targetName))
}

// func (gd GardenData) String() (outString string) {
// 	outString = fmt.Sprintf(
// 		"{ seeds: %v, xRefMap: [ %v ]}",
// 		gd.seeds,
// 		gd.xRefMap.String(),
// 	)
// 	return
// }
