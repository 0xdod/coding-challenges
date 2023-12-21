package main

import (
	"bufio"
	"bytes"
)

func countBytes(content []byte) int64 {
	return int64(len(content))
}

func countLines(content []byte) int64 {
	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	var count int64

	for scanner.Scan() {
		count++
	}

	return count
}

func countWords(content []byte) int64 {
	scanner := bufio.NewScanner(bytes.NewBuffer(content))
	scanner.Split(bufio.ScanWords)

	var count int64

	for scanner.Scan() {
		count++
	}

	return count
}

func countCharacters(content []byte) int64 {
	scanner := bufio.NewScanner(bytes.NewBuffer(content))

	scanner.Split(bufio.ScanRunes)

	var count int64

	for scanner.Scan() {
		count++
	}

	return count
}
