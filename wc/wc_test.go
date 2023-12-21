package main

import (
	"fmt"
	"os"
	"testing"
)

func getTestFileContent(t testing.TB, path string) []byte {
	t.Helper()

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

func Test_countBytes(t *testing.T) {
	input := getTestFileContent(t, "test.txt")
	want := int64(342190)
	got := countBytes(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test_countLines(t *testing.T) {
	input := getTestFileContent(t, "test.txt")
	want := int64(7145)
	got := countLines(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test_countWords(t *testing.T) {
	input := getTestFileContent(t, "test.txt")
	want := int64(58164)
	got := countWords(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test_countCharacters(t *testing.T) {
	input := getTestFileContent(t, "test.txt")
	want := int64(339292)
	got := countCharacters(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
