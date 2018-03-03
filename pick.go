package main

import (
	"regexp"
	"strings"
)

//********************************************************************
// Make pick list functions
//********************************************************************
// MakePickList is function.
// make simple string pick list.
func MakePickList(path string, defaltlist []string) ([]string, error) {
	picklist := make([]string, 0, len(defaltlist))
	if path != "" {
		var listErr error
		picklist, listErr = readFileLines(path)
		if listErr != nil {
			return picklist, listErr
		}
	}
	for _, pattern := range defaltlist {
		picklist = append(picklist, pattern)
	}
	return picklist, nil
}

// MakePickRegexpList is function.
// make regexp pattern pick list.
func MakePickRegexpList(picklist []string) ([]*regexp.Regexp, []error) {

	regexplist := make([]*regexp.Regexp, 0, len(picklist))
	errlist := make([]error, 0, len(picklist))
	for _, pattern := range picklist {
		re, reErr := regexp.Compile(pattern)
		if reErr != nil {
			errlist = append(errlist, reErr)
		} else {
			regexplist = append(regexplist, re)
		}
	}

	return regexplist, errlist
}

//********************************************************************
// Judge pick functions
//********************************************************************
// judgePick is function.
// if judge pick return true, else return false (contains word)
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

// judgePickRegexp is function.
// if judge pick return true, else return false (match regexp pattern)
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
