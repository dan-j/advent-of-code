package main

import (
	"bufio"
	"bytes"
	"os"
)

func Part2() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))
	fishes := readInput(s)

	// build a map of fishes of the same timer and a count of how many of them there are, since there are only 9 states
	// of fish (with timers from 0 through 9), we only have to process 9 fish on each day, just incrementing the counter
	// instead of building an array of size 1732821262171.
	fm := make(map[Fish]int)
	for _, f := range fishes {
		fm[f]++
	}

	for day := 1; day <= 256; day++ {
		next := make(map[Fish]int)
		for fish, count := range fm {
			if child, ok := fish.Decrement(); ok {
				next[child] += count
			}
			// you must += since there are fishes reproducing going from 0->6, and newborn fishes counting down from
			// 7->6
			next[fish] += count
		}

		fm = next
	}

	var total int
	for _, count := range fm {
		total += count
	}

	return total
}
