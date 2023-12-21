package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	var buffer []byte

	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			buffer = append(buffer, scanner.Bytes()...)
		}
		fmt.Println()
	} else {
		file, err := os.Open(files[0])

		if err != nil {
			fmt.Printf("ccwc: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		fi, _ := file.Stat()
		buffer = make([]byte, fi.Size())
		file.Read(buffer)
	}

	if linesCountFlag {
		linesCount = countLines(buffer)
		fmtStr = append(fmtStr, "%d ")
		fmtArgs = append(fmtArgs, linesCount)
	}

	if wordsCountFlag {
		wordsCount = countWords(buffer)
		fmtStr = append(fmtStr, "%d")
		fmtArgs = append(fmtArgs, wordsCount)
	}

	if charactersCountFlag {
		charsCount = countCharacters(buffer)
		fmtStr = append(fmtStr, "%d")
		fmtArgs = append(fmtArgs, charsCount)
	}

	if bytesCountFlag {
		bytesCount = countBytes(buffer)
		fmtStr = append(fmtStr, "%d")
		fmtArgs = append(fmtArgs, bytesCount)
	}

	if len(files) > 0 {
		fmtStr = append(fmtStr, "%s")
		fmtArgs = append(fmtArgs, files[0])
	}

	fmt.Printf("  "+strings.Join(fmtStr, " ")+"\n", fmtArgs...)
}
