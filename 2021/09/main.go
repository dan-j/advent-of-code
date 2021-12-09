package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	lines := parseInput(bufio.NewScanner(f))
	fmt.Println(Part1(lines))
	fmt.Println(Part2(lines))
}

func parseInput(s *bufio.Scanner) [][]int {
	var lines [][]int
	for s.Scan() {
		text := s.Text()
		line := make([]int, len(text))
		for i := range text {
			line[i] = atoi(string(text[i]))
		}

		lines = append(lines, line)
	}

	return lines
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
