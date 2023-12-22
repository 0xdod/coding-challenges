package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

func readFromStdIn() []byte {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
	var buffer []byte
	for scanner.Scan() {
		buffer = append(buffer, scanner.Bytes()...)
	}
	return buffer
}

func readFromFile(path string) []byte {
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("ccwc: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fi, _ := file.Stat()
	buffer := make([]byte, fi.Size())
	file.Read(buffer)

	return buffer
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		setDefaultFlags()
	}

	var (
		bytesCount int64
		linesCount int64
		wordsCount int64
		charsCount int64
	)

	var fmtStrs []string
	var fmtArgs []interface{}

	var buffer []byte

	args := flag.Args()

	if len(args) == 0 {
		buffer = readFromStdIn()
	} else {
		buffer = readFromFile(args[0])
	}

	if linesCountFlag {
		linesCount = countLines(buffer)
		fmtStrs = append(fmtStrs, "%d ")
		fmtArgs = append(fmtArgs, linesCount)
	}

	if wordsCountFlag {
		wordsCount = countWords(buffer)
		fmtStrs = append(fmtStrs, "%d")
		fmtArgs = append(fmtArgs, wordsCount)
	}

	if charactersCountFlag {
		charsCount = countCharacters(buffer)
		fmtStrs = append(fmtStrs, "%d")
		fmtArgs = append(fmtArgs, charsCount)
	}

	if bytesCountFlag {
		bytesCount = countBytes(buffer)
		fmtStrs = append(fmtStrs, "%d")
		fmtArgs = append(fmtArgs, bytesCount)
	}

	fmtStr := "  " + strings.Join(fmtStrs, " ")

	if len(args) > 0 {
		fmtStr = fmtStr + " %s\n"
		fmtArgs = append(fmtArgs, args[0])
	} else {
		fmtStr = "\n" + fmtStr + "\n"
	}

	fmt.Printf(fmtStr, fmtArgs...)
}
