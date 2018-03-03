package main

import (
	"errors"
	"flag"
	"fmt"
)

//********************************************************************
// Command arguments type
//********************************************************************
// CmdArgs is command arguments data object
type CmdArgs struct {
	srcPath    string   "srouce file path"
	listPath   string   "list file path"
	invFlag    bool     "pick judge invert flag"
	verFlag    bool     "show version flag"
	regexpFlag bool     "regexpFlag"
	length     int      "arguments  picklist length"
	picklist   []string "picklist from arguments"
}

//********************************************************************
// Command arguments methods
//********************************************************************
// Parse is CmdArgs method.
// command arguments parse. use flag package.
func (args *CmdArgs) Parse() {
	// options
	flag.StringVar(&args.srcPath, "s", "", "filter target text file.")
	flag.StringVar(&args.srcPath, "l", "", "match pattern list from list file and arguments.\nthe list file contents is 1 pattern per line.\nwhen not use `-s`, only command arguments.")
	flag.BoolVar(&args.regexpFlag, "regexp", false, "enable regexp pattern mode.\nwhen not use `-regexp`, match contains string line.")
	flag.BoolVar(&args.invFlag, "e", false, "enable exclude pattern match line mode.\nwhen not use `-e`, include pattern match line.")
	flag.BoolVar(&args.verFlag, "v", false, "show command version, and command exit.")

	flag.Parse()

	// arguments
	args.length = flag.NArg()
	args.picklist = flag.Args()
}

// DoCheckOption is CmdArgs method.
// check command options.
func (args *CmdArgs) DoCheckOption() (bool, error) {

	// show command
	if args.verFlag {
		fmt.Println(CMD_VERSION)
		// exit command, no error
		return true, nil
	}

	// check enable source file path
	if args.srcPath != "" && !doesExistPath(args.srcPath) {
		// exit command, error happen
		return true, errors.New(fmt.Sprintf("file not found :%s", args.srcPath))
	}

	// check enable list file path
	if args.listPath != "" && !doesExistPath(args.listPath) {
		// exit command, error happen
		return true, errors.New(fmt.Sprintf("file not found :%s", args.listPath))
	}

	// command proc continue
	return false, nil
}
