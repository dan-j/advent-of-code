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

	var inputs []int
	for s.Scan() {
		line := s.Text()
		x, _ := strconv.Atoi(line)

		inputs = append(inputs, x)
	}

	isFirst := true
	var prev int
	for i := 0; i < len(inputs)-2; i++ {
		x := sum(inputs[i : i+3]...)

		if isFirst {
			isFirst = false
			prev = x
			continue
		}

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
