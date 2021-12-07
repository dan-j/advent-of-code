package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, _ := os.ReadFile("test_input.txt")

	want := 37
	got := Part1(string(f))
	if got != want {
		t.Errorf("Part1(test_input) = %d, want = %d", got, want)
	}

	f, _ = os.ReadFile("input.txt")

	want = 344735
	got = Part1(string(f))
	if got != want {
		t.Errorf("Part1(input) = %d, want = %d", got, want)
	}
}
