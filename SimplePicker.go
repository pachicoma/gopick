package main

import (
	"strings"
)

//********************************************************************
// Data Defines
//********************************************************************
// simple text picker struct
type SimplePicker struct {
	pickList   []string
	invertFlag bool
	pickedItem string
}

//********************************************************************
// Picker Methods
//********************************************************************
// SetPickList return simple string pick list, read from file.
func (picker *SimplePicker) SetPickList(pickList []string) error {
	// return objects
	picker.pickList = pickList

	return nil
}

// SetInvert set pick judge result invert true or false
func (picker *SimplePicker) SetInvert(isInvert bool) {
	picker.invertFlag = isInvert
}

// JudgePick return pick or not pick judge result.
// when judged pick return true, else return false.
func (picker *SimplePicker) JudgePick(text string) bool {
	// return judge result value
	matched := false

	// do judge
	if len(picker.pickList) == 0 {
		matched = true
	} else {
		for _, pattern := range picker.pickList {
			if strings.Contains(text, pattern) {
				matched = true
				break
			}
		}
	}
	doesPick := (matched != picker.invertFlag)

	// pick item
	if doesPick {
		picker.pickedItem = text
	}

	return doesPick
}

// doPick return picked string
func (picker *SimplePicker) DoPick() string {
	return picker.pickedItem
}
