package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, _ := os.ReadFile("test_small.txt")

	want := 10 // short
	got := Part1(string(f))
	if got != want {
		t.Errorf("Part1(test_small) = %d, want = %d", got, want)
	}

	f, _ = os.ReadFile("test_medium.txt")
	want = 19
	got = Part1(string(f))
	if got != want {
		t.Errorf("Part1(test_medium) = %d, want = %d", got, want)
	}

	f, _ = os.ReadFile("input.txt")

	want = 3495
	got = Part1(string(f))
	if got != want {
		t.Errorf("Part1(input) = %d, want = %d", got, want)
	}
}
