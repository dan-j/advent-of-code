package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	f, _ := os.Open("test_input.txt")

	want := 61229
	got := Part2(parseInput(bufio.NewScanner(f)))
	if got != want {
		t.Errorf("Part1(test_input) = %d, want = %d", got, want)
	}

	f, _ = os.Open("input.txt")

	want = 1046281
	got = Part2(parseInput(bufio.NewScanner(f)))
	if got != want {
		t.Errorf("Part1(input) = %d, want = %d", got, want)
	}
}

func TestLine_Solve(t *testing.T) {
	l := &Line{
		Signals: strings.Split("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab", " "),
		Digits:  strings.Split("cdfeb fcadb cdfeb cdbaf", " "),
	}
	want := 5353
	if got := l.Solve(); got != want {
		t.Errorf("Solve() = %v, want %v", got, want)
	}
}
