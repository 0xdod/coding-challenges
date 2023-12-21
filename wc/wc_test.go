package main

import "testing"

func Test_countBytes(t *testing.T) {
	input := "test.txt"
	want := int64(342190)
	got := countBytes(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test_countLines(t *testing.T) {
	input := "test.txt"
	want := int64(7145)
	got := countLines(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test_countWords(t *testing.T) {
	input := "test.txt"
	want := int64(58164)
	got := countWords(input)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
