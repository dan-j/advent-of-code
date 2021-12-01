package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func Part1() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	var count int

	// scan the first line, it doesn't add to the count and is used at the initial `prev` value
	s.Scan()
	line := s.Text()
	prev, _ := strconv.Atoi(line)

	for s.Scan() {
		line = s.Text()
		x, _ := strconv.Atoi(line)

		if x > prev {
			count++
		}

		prev = x
	}

	return count
}
