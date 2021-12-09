package main

import (
	"bufio"
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
	f, _ := os.Open("test_input.txt")

	want := 1134
	got := Part2(parseInput(bufio.NewScanner(f)))
	if got != want {
		t.Errorf("Part1(test_input) = %d, want = %d", got, want)
	}

	f, _ = os.Open("input.txt")

	want = 882942
	got = Part2(parseInput(bufio.NewScanner(f)))
	if got != want {
		t.Errorf("Part1(input) = %d, want = %d", got, want)
	}
}
