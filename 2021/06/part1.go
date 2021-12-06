package main

import (
	"bufio"
	"bytes"
	"os"
)

func Part1() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	fishes := readInput(s)

	for day := 1; day <= 80; day++ {
		var spawned []Fish
		for i := range fishes {
			if child, ok := fishes[i].Decrement(); ok {
				spawned = append(spawned, child)
			}
		}

		fishes = append(fishes, spawned...)
	}

	return len(fishes)
}
