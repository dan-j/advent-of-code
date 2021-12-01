package main

import "testing"

func TestPart1(t *testing.T) {
	want := 1791
	got := Part1()
	if got != want {
		t.Errorf("Part2() = %d, want = %d", got, want)
	}
}
