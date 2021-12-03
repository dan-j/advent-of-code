package main

import "testing"

func TestPart2(t *testing.T) {
	want := 2073416724
	got := Part2()
	if got != want {
		t.Errorf("Part2() = %d, want = %d", got, want)
	}
}
