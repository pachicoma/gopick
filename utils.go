package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

const (
	ENC_SJIS      = "SJIS"
	ENC_UTF8      = "UTF8"
	ENC_EUC       = "EUCJP"
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
		reader = transform.NewReader(inStream, *encoding)
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
		writer = transform.NewWriter(outStream, *encoding)
	} else {
		writer = outStream
	}

	return writer
}

// getEncoding return encoding transformer object.
func getEncoding(enc string) *transform.Transformer {
	var encoding transform.Transformer
	switch strings.ToUpper(enc) {
	case ENC_SJIS:
		encoding = japanese.ShiftJIS.NewEncoder()
	case ENC_EUC:
		encoding = japanese.EUCJP.NewEncoder()
	case ENC_ISO2022JP:
		encoding = japanese.ISO2022JP.NewEncoder()
	default:
		return nil
	}
	return &encoding
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

	return lines, nil
}

// doesExistPath check path exist.
// when path exist return true, else return false.
func doesExistPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
