package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var (
		bytesCountFlag      bool
		linesCountFlag      bool
		wordsCountFlag      bool
		charactersCountFlag bool
	)

	flag.BoolVar(&bytesCountFlag, "c", false, "count bytes")
	flag.BoolVar(&linesCountFlag, "l", false, "count lines")
	flag.BoolVar(&wordsCountFlag, "w", false, "count words")
	flag.BoolVar(&charactersCountFlag, "m", false, "count characters")

	flag.Parse()

	files := flag.Args()
	// fmt.Println(files)

	if flag.NFlag() == 0 {
		linesCountFlag = true
		bytesCountFlag = true
		wordsCountFlag = true
	}

	var (
		bytesCount int64
		linesCount int64
		wordsCount int64
		charsCount int64
	)

	var fmtStr []string
	var fmtArgs []interface{}

	if linesCountFlag {
		linesCount = countLines(files[0])
		fmtStr = append(fmtStr, "%d ")
		fmtArgs = append(fmtArgs, linesCount)
	}

	if wordsCountFlag {
		wordsCount = countWords(files[0])
		fmtStr = append(fmtStr, "%d")
		fmtArgs = append(fmtArgs, wordsCount)
	}

	if charactersCountFlag {
		charsCount = countCharacters(files[0])
		fmtStr = append(fmtStr, "%d")
		fmtArgs = append(fmtArgs, charsCount)
	}

	if bytesCountFlag {
		bytesCount = countBytes(files[0])
		fmtStr = append(fmtStr, "%d")
		fmtArgs = append(fmtArgs, bytesCount)
	}

	fmtStr = append(fmtStr, "%s\n")
	fmtArgs = append(fmtArgs, files[0])

	fmt.Printf("  "+strings.Join(fmtStr, " "), fmtArgs...)
}
