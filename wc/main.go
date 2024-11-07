package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
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

func count(r io.Reader) (lineCount int64, wordCount int64, byteCount int64, charCount int64) {
	scanner := bufio.NewScanner(r)

	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// adapted the bufio.ScanLines implementation without the line endings dropped
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			return i + 1, data[0 : i+1], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	})

	for scanner.Scan() {
		line := scanner.Text()

		lineCount += 1

		wordCount += int64(len(strings.Fields(line)))
		charCount += int64(utf8.RuneCountInString(line))
		byteCount += int64(len(line))
	}

	return
}

func formatCount(count int64) string {
	return fmt.Sprintf("  %d", count)
}

func printCounts(lc, wc, cc, bc int64, path string) {
	text := ""

	if linesCountFlag {
		text += formatCount(lc)
	}

	if wordsCountFlag {
		text += formatCount(wc)
	}

	if charactersCountFlag {
		text += formatCount(cc)
	}

	if bytesCountFlag {
		text += formatCount(bc)
	}

	if path != "" {
		text += " " + path
	}

	fmt.Println(text)
}

func parseFlags() []string {
	flag.Parse()

	if flag.NFlag() == 0 {
		linesCountFlag = true
		bytesCountFlag = true
		wordsCountFlag = true
	}

	args := flag.Args()

	return args
}

func run(args []string) error {
	if len(args) == 0 {
		args = []string{""}
	}

	var printTotal bool

	if len(args) > 1 {
		printTotal = true
	}

	totalLc := int64(0)
	totalWc := int64(0)
	totalCc := int64(0)
	totalBc := int64(0)

	for _, arg := range args {
		path := arg
		var file *os.File

		if arg != "" {
			var err error

			file, err = os.Open(path)

			if err != nil {
				return err
			}

			defer file.Close()
		} else {
			file = os.Stdin
		}

		linesCount, wordsCount, bytesCount, charsCount := count(file)

		totalLc += linesCount
		totalWc += wordsCount
		totalCc += charsCount
		totalBc += bytesCount

		printCounts(linesCount, wordsCount, charsCount, bytesCount, path)
	}

	if printTotal {
		printCounts(totalLc, totalWc, totalCc, totalBc, "total")
	}

	return nil
}

func main() {
	args := parseFlags()

	if err := run(args); err != nil {
		fmt.Fprintf(os.Stderr, "ccwc: %v\n", err)
		os.Exit(1)
	}
}
