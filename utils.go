package main

import (
	"bufio"
	"os"
	"strings"
)

//********************************************************************
// File utility function
//********************************************************************
// readFileLines is function.
// read file lines convert to slice.
func readFileLines(path string) ([]string, error) {
	// return slice object
	lines := make([]string, 0, 0)

	// open file
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	// read pick list
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// skip blank line
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			continue
		}

		// add pick list
		lines = append(lines, line)
	}

	return lines, nil
}

// doesExistPath is function.
// if file exist at return true, else return false.
func doesExistPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
