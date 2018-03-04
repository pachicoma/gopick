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
	srcPath      string   "srouce file path"
	listPath     string   "list file path"
	invFlag      bool     "pick judge invert flag"
	verFlag      bool     "show version flag"
	rgxpFlag     bool     "regexp mode flag"
	length       int      "arguments  picklist length"
	picklist     []string "picklist from arguments"
	inEncoding   string   "input stream encoding"
	outEncoding  string   "output stream encoding"
	listEncoding string   "list file encoding"
}

// command options description
const (
	OPT_DESC_SRC  = `filter target text file.`
	OPT_DESC_LIST = `match pattern list from list file and arguments.
the list file contents is 1 pattern per line.
when not use "-s", only command arguments.`
	OPT_DESC_REGEXP = `enable regexp pattern mode.
when not use "-rgx", match contains string line.`
	OPT_DESC_I = `enable pattern unmatch(invert) line pick mode.
when not use "-i", include pattern match line.`
	OPT_DESC_V       = `show command version, and command exit.`
	OPT_DESC_ENCSRC  = `input stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
	OPT_DESC_ENCOUT  = `output stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
	OPT_DESC_ENCLIST = `list file encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
	OPT_DESC_ENC     = `list file and in/out stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
)

//********************************************************************
// Command arguments methods
//********************************************************************
// Parse is method of CmdArgs.
// command arguments parse and set default values.
// use flag package.
func (args *CmdArgs) Parse() {
	// options
	flag.StringVar(&args.srcPath, "src", "", OPT_DESC_SRC)
	flag.StringVar(&args.listPath, "list", "", OPT_DESC_LIST)
	flag.StringVar(&args.inEncoding, "encsrc", "", OPT_DESC_ENCSRC)
	flag.StringVar(&args.outEncoding, "encout", "", OPT_DESC_ENCOUT)
	flag.StringVar(&args.listEncoding, "enclist", "", OPT_DESC_ENCLIST)
	encoding := ""
	flag.StringVar(&encoding, "enc", "", OPT_DESC_ENC)
	flag.BoolVar(&args.rgxpFlag, "regexp", false, OPT_DESC_REGEXP)
	flag.BoolVar(&args.invFlag, "i", false, OPT_DESC_I)
	flag.BoolVar(&args.verFlag, "v", false, OPT_DESC_V)

	// parse arguments
	flag.Parse()

	// set default encoding
	args.inEncoding = selectStr(args.inEncoding, encoding, ENC_UTF8)
	args.outEncoding = selectStr(args.outEncoding, encoding, ENC_UTF8)
	args.listEncoding = selectStr(args.listEncoding, encoding, ENC_UTF8)

	// arguments
	args.length = flag.NArg()
	args.picklist = flag.Args()
}

// selectStr select from 3 strings if not "".
// priority first > second > defval
func selectStr(first, second, defval string) string {
	var selected string
	if first != "" {
		selected = first
	} else if second != "" {
		selected = second
	} else {
		selected = defval
	}
	return selected
}

// DoCheckOption is method of CmdArgs.
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
