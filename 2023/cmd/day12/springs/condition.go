package springs

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type (
	SpringCondition  string
	ConditionDetails struct {
		name    string
		matcher string
	}
)

const (
	SC_OPERATIONAL SpringCondition = "."
	SC_DAMAGED     SpringCondition = "#"
	SC_UNKNOWN     SpringCondition = "?"
)

var SpringConditions = map[SpringCondition]ConditionDetails{
	SC_OPERATIONAL: {"SC_OPERATIONAL", fmt.Sprintf("[%v%v]", string(SC_OPERATIONAL), string(SC_UNKNOWN))},
	SC_DAMAGED:     {"SC_DAMAGED", fmt.Sprintf("[%v%v]", string(SC_DAMAGED), string(SC_UNKNOWN))},
	SC_UNKNOWN:     {"SC_UNKNOWN", fmt.Sprintf("[%v]", string(SC_UNKNOWN))},
}

/*
**	METHODS
 */

func (sc SpringCondition) Name() (scName string) {
	if scd, ok := SpringConditions[sc]; ok {
		return scd.name
	} else {
		return string(sc)
	}
}

func (sc SpringCondition) GetSpringGroupMatchString(size int) string {
	return fmt.Sprintf("(%v){%v}", SpringConditions[sc].matcher, size)
}

/*
**	INTERFACES
 */

// (PipeType).String() string
func (sc SpringCondition) String() string {
	if scd, ok := SpringConditions[sc]; ok {
		return scd.name[len("SC_"):]
	} else {
		return string(sc)
	}
}

/*
**	HELPERS
 */

// StringToPipeType(pt) PipeType
func GetSpringCondition(scName string) SpringCondition {
	for _, sc := range maps.Keys(SpringConditions) {
		if scName == string(sc) || scName == sc.String() || scName == sc.Name() {
			return sc
		}
	}
	return SC_UNKNOWN
}
