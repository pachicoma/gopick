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
func MakePickList(path string, enc string, defaltlist []string) ([]string, error) {
	// return objects
	picklist := make([]string, 0, len(defaltlist))

	// make pick list
	if path != "" {
		var listErr error
		picklist, listErr = readNotBlankLines(path, enc)
		if listErr != nil {
			return picklist, listErr
		}
	}
	for _, pattern := range defaltlist {
		picklist = append(picklist, pattern)
	}
	return picklist, nil
}

// MakePickRegexpList return regexp pattern pick list, read from file.
func MakePickRegexpList(picklist []string) ([]*regexp.Regexp, []string) {
	// return objects
	regexplist := make([]*regexp.Regexp, 0, len(picklist))
	errmsgs := make([]string, 0, len(picklist))

	// make pick list
	for _, pattern := range picklist {
		re, reErr := regexp.Compile(pattern)
		if reErr != nil {
			errmsgs = append(errmsgs, fmt.Sprintf("\"%s\" > %s", pattern, reErr.Error()))
		} else {
			regexplist = append(regexplist, re)
		}
	}

	return regexplist, errmsgs
}

//********************************************************************
// Judge pick functions
//********************************************************************
// judgePick return pick or not pick judge result. (contains word mode)
// when judged pick return true, else return false.
func judgePick(text string, picklist []string, invertFlag bool) bool {
	// return judge result value
	doesPick := false

	// do judge
	if len(picklist) == 0 {
		doesPick = true
	} else {
		for _, pattern := range picklist {
			if strings.Contains(text, pattern) {
				doesPick = true
				break
			}
		}
	}

	return doesPick != invertFlag
}

// judgePickRegexp return pick or not pick judge result.(match regexp pattern mode)
// when judged pick return true, else return false.
func judgePickRegexp(text string, picklist []*regexp.Regexp, invertFlag bool) bool {
	// return judge result value
	doesPick := false

	// do judge
	if len(picklist) == 0 {
		doesPick = true
	} else {
		for _, pattern := range picklist {
			if pattern.MatchString(text) {
				doesPick = true
				break
			}
		}
	}

	return doesPick != invertFlag
}
