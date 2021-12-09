package main

import (
	"bufio"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, _ := os.Open("test_input.txt")

	want := 15
	got := Part1(parseInput(bufio.NewScanner(f)))
	if got != want {
		t.Errorf("Part1(test_input) = %d, want = %d", got, want)
	}

	f, _ = os.Open("input.txt")

	want = 558
	got = Part1(parseInput(bufio.NewScanner(f)))
	if got != want {
		t.Errorf("Part1(input) = %d, want = %d", got, want)
	}
}
