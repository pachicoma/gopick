package main

import (
	"errors"
	"flag"
	"fmt"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

//********************************************************************
// Command argument type
//********************************************************************
// CmdArgs is command argument data object
type CmdArgs struct {
	srcPath      string   "srouce file path"
	srcFiles     []string "srouce files path"
	listPath     string   "list file path"
	invFlag      bool     "pick judge invert flag"
	verFlag      bool     "show version flag"
	rgxpFlag     bool     "regexp mode flag"
	length       int      "argument  picklist length"
	pickList     []string "pick list by argument"
	inEncoding   string   "input stream encoding"
	outEncoding  string   "output stream encoding"
	listEncoding string   "pick list file encoding"
	lineStr      string   "range of file line(raw)"
	line         Range    "range of file line"
}

// Range start,end
type Range struct {
	start int
	end   int
}

// command options description
const (
	OPT_DESC_s = `filter target file.`

	OPT_DESC_l = `pick pattern list by file and argument.
the list file contents is 1 pattern per line.
when not use "-l", only command argument.`

	OPT_DESC_r = `pick range of target file line number.
you must give next format "-r start:end".
if start <= 0, pick start at first line.
if end > file max line number or end <= 0, pick to last of line.`

	OPT_DESC_regexp = `pick lines at regexp pattern matched.
when not use "-regexp", pick lines at contains string.`

	OPT_DESC_i = `pick lines at pattern unmatched.
when not use "-i", pick lines at pattern matched.`

	OPT_DESC_v = `show command version, and command exit.`

	OPT_DESC_es = `set to input stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
	OPT_DESC_eo = `set to output stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
	OPT_DESC_el = `set to list file encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
	OPT_DESC_e  = `set to list file and in/out stream encoding.(SJIS|EUCJP|ISO2022JP|UTF8)`
)

//********************************************************************
// Command argument methods
//********************************************************************
// Parse is method of CmdArgs.
// command argument parse and set default values.
// use flag package.
func (args *CmdArgs) Parse() {
	// options
	flag.StringVar(&args.srcPath, "s", "", OPT_DESC_s)
	flag.StringVar(&args.listPath, "l", "", OPT_DESC_l)
	flag.StringVar(&args.inEncoding, "es", "", OPT_DESC_es)
	flag.StringVar(&args.outEncoding, "eo", "", OPT_DESC_eo)
	flag.StringVar(&args.listEncoding, "el", "", OPT_DESC_el)
	flag.StringVar(&args.lineStr, "r", "0:0", OPT_DESC_r)
	encoding := ""
	flag.StringVar(&encoding, "e", "", OPT_DESC_e)
	flag.BoolVar(&args.rgxpFlag, "regexp", false, OPT_DESC_regexp)
	flag.BoolVar(&args.invFlag, "i", false, OPT_DESC_i)
	flag.BoolVar(&args.verFlag, "v", false, OPT_DESC_v)

	// parse argument
	flag.Parse()

	// set default encoding
	args.inEncoding = selectStr(args.inEncoding, encoding, ENC_UTF8)
	args.outEncoding = selectStr(args.outEncoding, encoding, ENC_UTF8)
	args.listEncoding = selectStr(args.listEncoding, encoding, ENC_UTF8)

	// othrer argument
	args.length = flag.NArg()
	args.pickList = flag.Args()
}

// DoCheckOption is method of CmdArgs.
// check command options.
func (args *CmdArgs) DoCheckOptions() (bool, error) {

	// show command
	if args.verFlag {
		fmt.Println(CMD_VERSION)
		// exit command, no error
		return true, nil
	}

	// set target range
	var rangeErr error
	lineNums := strings.Split(args.lineStr, ":")
	if len(lineNums) != 2 {
		return true, errors.New(fmt.Sprintf("give bad range format: '%s', you must give format 'start:end'", args.lineStr))
	}
	args.line.start, rangeErr = convertLineNumberToInt(lineNums[0], 0)
	if rangeErr != nil {
		return true, rangeErr
	}
	args.line.end, rangeErr = convertLineNumberToInt(lineNums[1], 0)
	if rangeErr != nil {
		return true, rangeErr
	}
	if args.line.end > 0 && args.line.start > args.line.end {
		return true, errors.New(fmt.Sprintf("give bad range value: '%s', you must give value start < end", args.lineStr))
	}

	// check enable source file path
	if args.srcPath != "" {
		// use -s option, read from file
		args.srcPath = filepath.Clean(args.srcPath)
		// '~' convert home directory
		if args.srcPath[0] == '~' {
			usr, errUsr := user.Current()
			if errUsr != nil {
				return true, errors.New(fmt.Sprintf("error of get user info: %s", errUsr.Error()))
			}
			args.srcPath = strings.Replace(args.srcPath, "~", usr.HomeDir, 1)
		}
		// wildcard analyze
		files, srcErr := filepath.Glob(args.srcPath)
		if srcErr != nil {
			return true, errors.New(fmt.Sprintf("give bad source file path :%s", args.srcPath))
		}
		args.srcFiles = make([]string, 0, 1)
		for _, file := range files {
			// convert to absolute path
			absPath, errPath := filepath.Abs(file)
			if errPath != nil {
				return true, errors.New(fmt.Sprintf("'%s': error of abs path: %s", file, errPath.Error()))
			}
			args.srcFiles = append(args.srcFiles, absPath)
		}
	} else {
		// not use -s, read from stdin
		args.srcFiles = []string{""}
	}

	// check enable list file path
	if args.listPath != "" && !DoesExistPath(args.listPath) {
		return true, errors.New(fmt.Sprintf("pick list file not found :%s", args.listPath))
	}

	// command proc continue
	return false, nil
}

//********************************************************************
// Helper functions
//********************************************************************
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

// getRangeLineNumber return parse to string range of line number
func convertLineNumberToInt(numStr string, min int) (int, error) {
	var lineNum int
	if numStr == "" {
		lineNum = min
	} else {
		lineNum64, rangeErr := strconv.ParseInt(numStr, 10, 0)
		if rangeErr != nil {
			return min, errors.New(fmt.Sprintf("give bad value: '%s'", numStr))
		}
		lineNum = int(lineNum64)
		if lineNum < min {
			lineNum = min
		}
	}
	return lineNum, nil
}
