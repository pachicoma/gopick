//******************************************************************************
// `gopick` is simple filter command.
// If arguments pattern matches each line of input text data,
// then match line text data to the stdout stream.
// You can select mode, matched line exclude or include mode.
// You can select input source, file or stdin stream.
// You can select match pattern list srouce, file or command arguments.
//
// USAGE:
//   $gopick [-s srcfile] [-l listfile] [-regexp] [-e] [-v] pattern
//
//     -s srcfile   filter target text file.
//                  when not use `-s`, read of stdin stream.
//
//     -l listfile  match pattern list from list file and arguments.
//                  the list file contents is 1 pattern per line.
//                  when not use `-l`, only command arguments.
//
//     -regexp      enable regexp pattern mode.
//                  when not use `-regexp`, match contains string line.
//
//     -e           enable exclude pattern match line mode.
//                  when not use `-e`, include pattern match line.
//
//     -v           show command version, and command exit.
//
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
	picklist, listErr := MakePickList(args.listPath, args.picklist)
	if listErr != nil {
		log.Fatalln(listErr.Error())
	}
	if args.regexpFlag {
		// if use regexp pattern mode
		var regexpErrs []error
		regexplist, regexpErrs = MakePickRegexpList(picklist)
		for _, regexpErr := range regexpErrs {
			log.Println(regexpErr.Error())
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
		scanner = bufio.NewScanner(srcFile)

	} else {
		// from stdin
		scanner = bufio.NewScanner(os.Stdin)
	}

	// do pick and output
	if args.regexpFlag {
		// match regexp pattern
		for scanner.Scan() {
			lineText := scanner.Text()
			if judgePickRegexp(lineText, regexplist, args.invFlag) {
				fmt.Println(lineText)
			}
		}
	} else {
		// contains string
		for scanner.Scan() {
			lineText := scanner.Text()
			if judgePick(lineText, picklist, args.invFlag) {
				fmt.Println(lineText)
			}
		}
	}
}
