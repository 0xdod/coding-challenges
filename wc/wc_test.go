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
