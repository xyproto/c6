// generated by stringer -type=UnitType token.go unit.go; DO NOT EDIT

package ast

import "fmt"

const _UnitType_name = "UNIT_NONEUNIT_PXUNIT_PTUNIT_EMUNIT_CMUNIT_MMUNIT_REMUNIT_DEGUNIT_PERCENTUNIT_SECONDUNIT_MILLISECOND"

var _UnitType_index = [...]uint8{0, 9, 16, 23, 30, 37, 44, 52, 60, 72, 83, 99}

func (i UnitType) String() string {
	if i < 0 || i+1 >= UnitType(len(_UnitType_index)) {
		return fmt.Sprintf("UnitType(%d)", i)
	}
	return _UnitType_name[_UnitType_index[i]:_UnitType_index[i+1]]
}
