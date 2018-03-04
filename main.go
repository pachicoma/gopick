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
const CMD_VERSION = "0.02"

// entry point function
func main() {
	var args CmdArgs
	var pickList []string
	var rgxpList []*regexp.Regexp

	// analyze command arguments
	args.Parse()
	exit, optErr := args.DoCheckOptions()
	if exit {
		if optErr != nil {
			log.Fatalln(optErr.Error())
		}
		os.Exit(0)
	}

	// make pick list
	pickList, listErr := MakePickList(args.listPath, args.listEncoding, args.pickList)
	if listErr != nil {
		log.Fatalln(fmt.Sprintf("makelist file read error: %s", listErr.Error()))
	}
	if args.rgxpFlag {
		// if use regexp pattern mode
		var regexpErrMsgs []string
		rgxpList, regexpErrMsgs = MakePickRegexpList(pickList)
		for _, errMsg := range regexpErrMsgs {
			log.Println(errMsg)
		}
	}

	// select input source
	var scanner *bufio.Scanner
	writer := bufio.NewWriter(NewEncodeWriter(os.Stdout, args.outEncoding))
	for _, srcFilePath := range args.srcFiles {
		var srcFile *os.File
		if srcFilePath != "" {
			// from file
			var srcErr error
			srcFile, srcErr = os.Open(srcFilePath)
			if srcErr != nil {
				log.Fatalln(srcErr.Error())
			}

			scanner = bufio.NewScanner(NewDecodeReader(srcFile, args.inEncoding))
		} else {
			// from stdin
			scanner = bufio.NewScanner(NewDecodeReader(os.Stdin, args.inEncoding))
		}

		// skip for start of pick range
		var lineNum int
		for lineNum = 1; lineNum < args.line.start; lineNum++ {
			scanner.Scan()
		}

		// do pick and output
		if args.rgxpFlag {
			// match regexp pattern
			for scanner.Scan() {
				lineText := scanner.Text()
				if JudgePickRegexp(lineText, rgxpList, args.invFlag) {
					fmt.Fprintln(writer, lineText)
				}
				// check end of pick range
				if args.line.end > 0 {
					if lineNum >= args.line.end {
						break
					}
					lineNum++
				}
			}
		} else {
			// contains string
			for scanner.Scan() {
				lineText := scanner.Text()
				if JudgePick(lineText, pickList, args.invFlag) {
					fmt.Fprintln(writer, lineText)
				}
				// check end of pick range
				if args.line.end > 0 {
					if lineNum >= args.line.end {
						break
					}
					lineNum++
				}
			}
		}
		if err := scanner.Err(); err != nil {
			log.Printf("taget file read error: %s", err.Error())
		}

		// close file
		srcFile.Close()
	}
	writer.Flush()
}
