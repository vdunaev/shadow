// Code generated by "enumer -type=componentStatus -trimprefix=ComponentStatus -output=component_status_enumer.go -transform=snake /Users/kihamo/go/src/github.com/kihamo/shadow"; DO NOT EDIT.

package shadow

import (
	"fmt"
)

const _componentStatusName = "unknownreadyrun_failedfinishedshutdown"

var _componentStatusIndex = [...]uint8{0, 7, 12, 22, 30, 38}

func (i componentStatus) String() string {
	if i < 0 || i >= componentStatus(len(_componentStatusIndex)-1) {
		return fmt.Sprintf("componentStatus(%d)", i)
	}
	return _componentStatusName[_componentStatusIndex[i]:_componentStatusIndex[i+1]]
}

var _componentStatusValues = []componentStatus{0, 1, 2, 3, 4}

var _componentStatusNameToValueMap = map[string]componentStatus{
	_componentStatusName[0:7]:   0,
	_componentStatusName[7:12]:  1,
	_componentStatusName[12:22]: 2,
	_componentStatusName[22:30]: 3,
	_componentStatusName[30:38]: 4,
}

// componentStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func componentStatusString(s string) (componentStatus, error) {
	if val, ok := _componentStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to componentStatus values", s)
}

// componentStatusValues returns all values of the enum
func componentStatusValues() []componentStatus {
	return _componentStatusValues
}

// IsAcomponentStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i componentStatus) IsAcomponentStatus() bool {
	for _, v := range _componentStatusValues {
		if i == v {
			return true
		}
	}
	return false
}