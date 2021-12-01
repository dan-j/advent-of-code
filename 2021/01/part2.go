package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

// Part2 has the answer:  1822
func Part2() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	var count int
	lines := make([]int, 0, 3)

	// read the first 3 lines
	for i := 0; i < 3 && s.Scan(); i++ {
		num, _ := strconv.Atoi(s.Text())
		lines = append(lines, num)
	}

	if len(lines) < 3 {
		panic("couldn't read 3 lines from input")
	}

	prev := sum(lines...)
	for s.Scan() {
		copy(lines, lines[1:])
		num, _ := strconv.Atoi(s.Text())

		lines[2] = num
		x := sum(lines...)

		if x > prev {
			count++
		}

		prev = x
	}

	return count
}

func sum(vals ...int) (r int) {
	for _, v := range vals {
		r += v
	}

	return r
}
