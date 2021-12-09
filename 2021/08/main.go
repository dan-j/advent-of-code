package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var digitDefinitions = map[string]Set{
	"0": NewSet("abcefg"),
	"1": NewSet("cf"),
	"2": NewSet("acdeg"),
	"3": NewSet("acdfg"),
	"4": NewSet("bcdf"),
	"5": NewSet("abdfg"),
	"6": NewSet("abdefg"),
	"7": NewSet("acf"),
	"8": NewSet("abcdefg"),
	"9": NewSet("abcdfg"),
}

func main() {
	f, _ := os.Open("input.txt")
	lines := parseInput(bufio.NewScanner(f))
	fmt.Println(Part1(lines))
	fmt.Println(Part2(lines))
}

type Line struct {
	Signals []string
	Digits  []string
}

func parseInput(s *bufio.Scanner) []Line {
	var lines []Line
	for s.Scan() {
		split := strings.Split(s.Text(), "|")
		signalInput, digitsInput := strings.TrimSpace(split[0]), strings.TrimSpace(split[1])

		lines = append(lines, Line{
			Signals: strings.Split(signalInput, " "),
			Digits:  strings.Split(digitsInput, " "),
		})
	}

	return lines
}
