package main

import (
	"bufio"
	"strconv"
	"strings"
)

func readInput(s *bufio.Scanner) Fishes {
	s.Scan()
	timers := strings.Split(s.Text(), ",")
	fishes := make(Fishes, len(timers))
	for i := range timers {
		fishes[i] = Fish{Timer: atoi(timers[i])}
	}

	return fishes
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
