package main

import (
	"bufio"
	"os"
)

func countBytes(filePath string) int64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fi, _ := file.Stat()

	return fi.Size()
}

func countLines(filePath string) int64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int64

	for scanner.Scan() {
		count++
	}

	return count
}

func countWords(filePath string) int64 {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var count int64

	for scanner.Scan() {
		count++
	}

	return count
}

func countCharacters(filePath string) int64 {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	var count int64

	for scanner.Scan() {
		count++
	}

	return count
}
