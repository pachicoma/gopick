//******************************************************************************
// `gopick` is simple filter command.
//******************************************************************************
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// command version
const CMD_VERSION = "0.01"

// entry point function
func main() {
	var args CmdArgs
	var picklist []string
	var regexplist []*regexp.Regexp

	// analyze command arguments
	args.Parse()
	exit, optErr := args.DoCheckOption()
	if exit {
		if optErr != nil {
			log.Fatalln(optErr.Error())
		}
		os.Exit(0)
	}

	// make pick list
	picklist, listErr := MakePickList(args.listPath, args.listEncoding, args.picklist)
	if listErr != nil {
		log.Fatalln(listErr.Error())
	}
	if args.rgxpFlag {
		// if use regexp pattern mode
		var regexpErrMsgs []string
		regexplist, regexpErrMsgs = MakePickRegexpList(picklist)
		for _, errMsg := range regexpErrMsgs {
			log.Println(errMsg)
		}
	}

	// select input source
	var scanner *bufio.Scanner
	if args.srcPath != "" {
		// from file
		srcFile, srcErr := os.Open(args.srcPath)
		if srcErr != nil {
			log.Fatalln(srcErr.Error())
		}
		defer srcFile.Close()

		scanner = bufio.NewScanner(getTextReader(srcFile, args.inEncoding))
	} else {
		// from stdin
		scanner = bufio.NewScanner(getTextReader(os.Stdin, args.inEncoding))
	}

	// do pick and output
	writer := getTextWriter(os.Stdout, args.outEncoding)
	if args.rgxpFlag {
		// match regexp pattern
		for scanner.Scan() {
			lineText := scanner.Text()
			if judgePickRegexp(lineText, regexplist, args.invFlag) {
				fmt.Fprintln(writer, lineText)
			}
		}
	} else {
		// contains string
		for scanner.Scan() {
			lineText := scanner.Text()
			if judgePick(lineText, picklist, args.invFlag) {
				fmt.Fprintln(writer, lineText)
			}
		}
		err := scanner.Err()
		if err != nil {
			log.Printf("scan err: %s", err.Error())
		}
	}
}
