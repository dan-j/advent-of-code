package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

func Part1() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	p := &Part1Position{}

	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "forward":
			p.Forward(atoi(parts[1]))
		case "down":
			p.Down(atoi(parts[1]))
		case "up":
			p.Up(atoi(parts[1]))
		}
	}

	return p.Result()
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

type Part1Position struct {
	forward int
	depth   int
}

func (p *Part1Position) Forward(n int) {
	p.forward += n
}

func (p *Part1Position) Down(n int) {
	p.depth += n
}

func (p *Part1Position) Up(n int) {
	p.depth -= n
}

func (p *Part1Position) Result() int {
	return p.forward * p.depth
}
