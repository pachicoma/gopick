package main

import (
	"fmt"
	"regexp"
	"strings"
)

//********************************************************************
// Make pick list functions
//********************************************************************
// MakePickList return simple string pick list, read from file.
func MakePickList(path string, enc string, defaltPickList []string) ([]string, error) {
	// return objects
	pickList := make([]string, 0, len(defaltPickList))

	// make pick list
	if path != "" {
		var listErr error
		pickList, listErr = ReadNotBlankLines(path, enc)
		if listErr != nil {
			return pickList, listErr
		}
	}
	// append default pick list
	for _, pattern := range defaltPickList {
		pickList = append(pickList, pattern)
	}
	return pickList, nil
}

// MakePickRegexpList return regexp pattern pick list, read from file.
func MakePickRegexpList(pickList []string) ([]*regexp.Regexp, []string) {
	// return objects
	rgxpList := make([]*regexp.Regexp, 0, len(pickList))
	errMsgs := make([]string, 0, len(pickList))

	// make pick list
	for _, pattern := range pickList {
		re, reErr := regexp.Compile(pattern)
		if reErr != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("\"%s\" > %s", pattern, reErr.Error()))
		} else {
			rgxpList = append(rgxpList, re)
		}
	}

	return rgxpList, errMsgs
}

//********************************************************************
// Judge pick functions
//********************************************************************
// JudgePick return pick or not pick judge result. (contains word mode)
// when judged pick return true, else return false.
func JudgePick(text string, pickList []string, invertFlag bool) bool {
	// return judge result value
	doesPick := false

	// do judge
	if len(pickList) == 0 {
		doesPick = true
	} else {
		for _, pattern := range pickList {
			if strings.Contains(text, pattern) {
				doesPick = true
				break
			}
		}
	}

	return doesPick != invertFlag
}

// JudgePickRegexp return pick or not pick judge result.(match regexp pattern mode)
// when judged pick return true, else return false.
func JudgePickRegexp(text string, pickList []*regexp.Regexp, invertFlag bool) bool {
	// return judge result value
	doesPick := false

	// do judge
	if len(pickList) == 0 {
		doesPick = true
	} else {
		for _, pattern := range pickList {
			if pattern.MatchString(text) {
				doesPick = true
				break
			}
		}
	}

	return doesPick != invertFlag
}
