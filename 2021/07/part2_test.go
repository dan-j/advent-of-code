package main

import (
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
	f, _ := os.ReadFile("test_input.txt")

	want := 168
	got := Part2(string(f))
	if got != want {
		t.Errorf("Part2(test_input) = %d, want = %d", got, want)
	}

	f, _ = os.ReadFile("input.txt")

	want = 96798233
	got = Part2(string(f))
	if got != want {
		t.Errorf("Part2(input) = %d, want = %d", got, want)
	}
}
