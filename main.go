//******************************************************************************
// `gopick` is simple filter command.
//******************************************************************************
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// command version
const CMD_VERSION = "0.03"

//********************************************************************
// entry point function
//********************************************************************
func main() {
	var args CmdArgs
	var pickList []string
	var picker Picker

	// analyze command arguments
	args.Parse()
	exit, optErr := args.DoCheckOptions()
	if exit {
		JudgeFatalError(optErr, "arguments error")
		os.Exit(0)
	}

	// make picklist
	pickList, listErr := MakePickListFromFile(args.listPath, args.listEncoding)
	JudgeFatalError(listErr, "read pick list file error")
	pickList = append(pickList, args.pickList...)

	// make picker
	if args.rgxpFlag {
		picker = &RegexpPicker{}
	} else {
		picker = &SimplePicker{}
	}
	pickerErr := picker.SetPickList(pickList)
	JudgeFatalError(pickerErr, "make pick list error")
	picker.SetInvert(args.invFlag)

	// select input source
	writer := bufio.NewWriter(NewEncodeWriter(os.Stdout, args.outEncoding))
	for _, srcFilePath := range args.srcFiles {

		// open input stream and make scanner
		srcFile, scanner, scanErr := MakeFileScaner(srcFilePath, args.inEncoding)
		JudgeMinorError(scanErr, "input stream open error")

		// skip for start of pick range
		var lineNum int
		for lineNum = 1; lineNum < args.line.start; lineNum++ {
			scanner.Scan()
		}

		// do pick and output picked data
		if args.line.end > 0 {
			// pick for end of range
			for scanner.Scan() {
				if picker.JudgePick(scanner.Text()) {
					fmt.Fprintln(writer, picker.DoPick())
				}
				// check end of pick range
				if lineNum >= args.line.end {
					break
				}
				lineNum++
			}
		} else {
			// pick for EOF
			for scanner.Scan() {
				if picker.JudgePick(scanner.Text()) {
					fmt.Fprintln(writer, picker.DoPick())
				}
			}
		}
		JudgeMinorError(scanner.Err(), "input stream scan error")

		// close file
		srcFile.Close()
	}
	// flush buffer data to output stream
	writer.Flush()
}

//********************************************************************
// Helper functions
//********************************************************************
// MakeFileScaner make scanner from file or stdin stream
func MakeFileScaner(filePath string, encoding string) (*os.File, *bufio.Scanner, error) {
	var file *os.File
	var scanner *bufio.Scanner
	if filePath != "" {
		// from file
		var err error
		file, err = os.Open(filePath)
		if err != nil {
			return nil, nil, err
		}

		scanner = bufio.NewScanner(NewDecodeReader(file, encoding))
	} else {
		// from stdin
		scanner = bufio.NewScanner(NewDecodeReader(os.Stdin, encoding))
	}

	return file, scanner, nil
}

// JudgeFatalError
func JudgeFatalError(err error, msg string) {
	if err != nil {
		log.Fatalln(fmt.Sprintf("%s: %s", msg, err.Error()))
	}
}

// JudgeMinorError
func JudgeMinorError(err error, msg string) {
	if err != nil {
		log.Println(fmt.Sprintf("%s: %s", msg, err.Error()))
	}
}
