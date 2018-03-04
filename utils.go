package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	ENC_SJIS      = "SJIS"
	ENC_SHIFTJIS  = "SHIFTJIS"
	ENC_SHIFT_JIS = "SHIFT_JIS"
	ENC_UTF8      = "UTF8"
	ENC_UTF_8     = "UTF-8"
	ENC_EUCJP     = "EUCJP"
	ENC_EUC_JP    = "EUC-JP"
	ENC_ISO2022JP = "ISO2022JP"
)

//********************************************************************
// Charset encoding functions
//********************************************************************
// getTextReader return io.Reader object
func getTextReader(inStream io.Reader, enc string) io.Reader {
	var reader io.Reader

	encoding := getEncoding(enc)
	if encoding != nil {
		reader = transform.NewReader(inStream, encoding.NewDecoder())
	} else {
		reader = inStream
	}

	return reader
}

// getTextWriter return io.Writer object
func getTextWriter(outStream io.Writer, enc string) io.Writer {
	var writer io.Writer

	encoding := getEncoding(enc)
	if encoding != nil {
		writer = transform.NewWriter(outStream, encoding.NewEncoder())
	} else {
		writer = outStream
	}

	return writer
}

// getEncoding return encoding transformer object.
func getEncoding(enc string) encoding.Encoding {
	switch strings.ToUpper(enc) {
	case ENC_SJIS, ENC_SHIFTJIS, ENC_SHIFT_JIS:
		return japanese.ShiftJIS
	case ENC_EUCJP, ENC_EUC_JP:
		return japanese.EUCJP
	case ENC_ISO2022JP:
		return japanese.ISO2022JP
	default:
		return nil
	}
}

//********************************************************************
// File utility function
//********************************************************************
// readNotBlankLines convert text file lines to slice, only not blank line
func readNotBlankLines(path string, enc string) ([]string, error) {
	// return slice object
	lines := make([]string, 0, 0)

	// open file
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	// read pick list
	scanner := bufio.NewScanner(getTextReader(file, enc))
	for scanner.Scan() {
		// skip blank line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// add pick list
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}

// doesExistPath check path exist.
// when path exist return true, else return false.
func doesExistPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
