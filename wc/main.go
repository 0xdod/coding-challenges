package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var (
		bytesCountFlag bool
		linesCountFlag bool
		wordsCountFlag bool
	)

	flag.BoolVar(&bytesCountFlag, "c", false, "count bytes")
	flag.BoolVar(&linesCountFlag, "l", false, "count lines")
	flag.BoolVar(&wordsCountFlag, "w", false, "count words")

	flag.Parse()

	files := flag.Args()
	fmt.Println(files)

	var (
		bytesCount int64
		linesCount int64
		wordsCount int64
	)

	var fmtStr []string
	var fmtArgs []interface{}

	if linesCountFlag {
		linesCount = countLines(files[0])
		fmtStr = append(fmtStr, "lines: %d")
		fmtArgs = append(fmtArgs, linesCount)
	}

	if wordsCountFlag {
		wordsCount = countWords(files[0])
		fmtStr = append(fmtStr, "words: %d")
		fmtArgs = append(fmtArgs, wordsCount)
	}

	if bytesCountFlag {
		bytesCount = countBytes(files[0])
		fmtStr = append(fmtStr, "bytes: %d")
		fmtArgs = append(fmtArgs, bytesCount)
	}

	fmt.Printf(strings.Join(fmtStr, " ")+"\n", fmtArgs...)
}
