package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var (
	bytesCountFlag      bool
	linesCountFlag      bool
	wordsCountFlag      bool
	charactersCountFlag bool
)

func init() {
	flag.BoolVar(&bytesCountFlag, "c", false, "count bytes")
	flag.BoolVar(&linesCountFlag, "l", false, "count lines")
	flag.BoolVar(&wordsCountFlag, "w", false, "count words")
	flag.BoolVar(&charactersCountFlag, "m", false, "count characters")
}

func setDefaultFlags() {
	linesCountFlag = true
	bytesCountFlag = true
	wordsCountFlag = true
}

func wc() {
	var (
		bytesCount int64
		linesCount int64
		wordsCount int64
		charsCount int64
	)

	var fmtStrs []string
	var fmtArgs []interface{}

	args := flag.Args()

	var file *os.File

	if len(args) == 0 {
		file = os.Stdin
	} else {
		var err error
		path := args[0]

		file, err = os.Open(path)

		if err != nil {
			fmt.Printf("ccwc: %v\n", err)
			os.Exit(1)
		}

		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)

		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			if len(data[:i]) > 0 && data[i-1] == '\r' {
				token = append(token, '\r')
			}

			token = append(token, '\n')
		}

		return
	})

	for scanner.Scan() {
		line := scanner.Text()

		linesCount += 1

		wordsCount += int64(len(strings.Fields(line)))
		charsCount += int64(utf8.RuneCountInString(line))
		bytesCount += int64(len(line))
	}

	if linesCountFlag {
		fmtStrs = append(fmtStrs, "%d ")
		fmtArgs = append(fmtArgs, linesCount)
	}

	if wordsCountFlag {
		fmtStrs = append(fmtStrs, "%d")
		fmtArgs = append(fmtArgs, wordsCount)
	}

	if charactersCountFlag {
		fmtStrs = append(fmtStrs, "%d")
		fmtArgs = append(fmtArgs, charsCount)
	}

	if bytesCountFlag {
		fmtStrs = append(fmtStrs, "%d")
		fmtArgs = append(fmtArgs, bytesCount)
	}

	fmtStr := "  " + strings.Join(fmtStrs, " ")

	if len(args) > 0 {
		fmtStr = fmtStr + " %s\n"
		fmtArgs = append(fmtArgs, args[0])
	} else {
		fmtStr = fmtStr + "\n"
	}

	fmt.Printf(fmtStr, fmtArgs...)
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		setDefaultFlags()
	}

	wc()
}
