package main

//********************************************************************
// Interface
//********************************************************************
// Picker interface
type Picker interface {
	SetPickList(pickList []string) error
	SetInvert(isInvert bool)
	JudgePick(text string) bool
	DoPick() string
}

//********************************************************************
// Helper functions
//********************************************************************
// MakePickListFromFile return text pick list([]string), read from file.
func MakePickListFromFile(path string, enc string) ([]string, error) {
	// return objects
	pickList := make([]string, 0, 0)

	// make pick list
	if path != "" {
		var listErr error
		pickList, listErr = ReadNotBlankLines(path, enc)
		if listErr != nil {
			return pickList, listErr
		}
	}

	return pickList, nil
}
