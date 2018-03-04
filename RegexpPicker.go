package main

import (
	"regexp"
)

//********************************************************************
// Data Defines
//********************************************************************
// regexp text picker struct
type RegexpPicker struct {
	pickList   []*regexp.Regexp
	invertFlag bool
	pickedItem string
}

//********************************************************************
// Picker Methods
//********************************************************************
// SetPickList return regexp pattern pick list
func (picker *RegexpPicker) SetPickList(pickList []string) error {
	// return objects
	picker.pickList = make([]*regexp.Regexp, 0, len(pickList))

	// make pick list
	for _, pattern := range pickList {
		reObj, reErr := regexp.Compile(pattern)
		if reErr != nil {
			return reErr
		} else {
			picker.pickList = append(picker.pickList, reObj)
		}
	}

	return nil
}

// SetInvert set pick judge result invert true or false
func (picker *RegexpPicker) SetInvert(isInvert bool) {
	picker.invertFlag = isInvert
}

// JudgePick return pick or not pick judge result.
// when judged pick return true, else return false.
func (picker *RegexpPicker) JudgePick(text string) bool {
	// return judge result value
	matched := false

	// do judge
	if len(picker.pickList) == 0 {
		matched = true
	} else {
		for _, pattern := range picker.pickList {
			if pattern.MatchString(text) {
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
func (picker *RegexpPicker) DoPick() string {
	return picker.pickedItem
}
