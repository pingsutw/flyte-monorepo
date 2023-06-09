// Code generated by "enumer --type=EPhase --trimprefix=EPhase"; DO NOT EDIT.

package handler

import (
	"fmt"
)

const _EPhaseName = "UndefinedNotReadyQueuedRunningSkipFailedRetryableFailureSuccessTimedoutFailingDynamicRunningRecovered"

var _EPhaseIndex = [...]uint8{0, 9, 17, 23, 30, 34, 40, 56, 63, 71, 78, 92, 101}

func (i EPhase) String() string {
	if i >= EPhase(len(_EPhaseIndex)-1) {
		return fmt.Sprintf("EPhase(%d)", i)
	}
	return _EPhaseName[_EPhaseIndex[i]:_EPhaseIndex[i+1]]
}

var _EPhaseValues = []EPhase{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

var _EPhaseNameToValueMap = map[string]EPhase{
	_EPhaseName[0:9]:    0,
	_EPhaseName[9:17]:   1,
	_EPhaseName[17:23]:  2,
	_EPhaseName[23:30]:  3,
	_EPhaseName[30:34]:  4,
	_EPhaseName[34:40]:  5,
	_EPhaseName[40:56]:  6,
	_EPhaseName[56:63]:  7,
	_EPhaseName[63:71]:  8,
	_EPhaseName[71:78]:  9,
	_EPhaseName[78:92]:  10,
	_EPhaseName[92:101]: 11,
}

// EPhaseString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EPhaseString(s string) (EPhase, error) {
	if val, ok := _EPhaseNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EPhase values", s)
}

// EPhaseValues returns all values of the enum
func EPhaseValues() []EPhase {
	return _EPhaseValues
}

// IsAEPhase returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EPhase) IsAEPhase() bool {
	for _, v := range _EPhaseValues {
		if i == v {
			return true
		}
	}
	return false
}
