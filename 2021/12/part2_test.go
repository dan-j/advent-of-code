package main

import (
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
	f, _ := os.ReadFile("test_small.txt")

	want := 36
	got := Part2(string(f))
	if got != want {
		t.Errorf("Part1(test_input) = %d, want = %d", got, want)
	}

	f, _ = os.ReadFile("test_medium.txt")

	want = 103
	got = Part2(string(f))
	if got != want {
		t.Errorf("Part1(test_medium) = %d, want = %d", got, want)
	}

	f, _ = os.ReadFile("input.txt")

	want = 0
	got = Part2(string(f))
	if got != want {
		t.Errorf("Part1(input) = %d, want = %d", got, want)
	}
}
