package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func Part2() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	p := &Part2Position{}

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

type Part2Position struct {
	aim int
	x   int
	z   int
}

func (p *Part2Position) Up(n int) {
	p.aim -= n
}

func (p *Part2Position) Down(n int) {
	p.aim += n
}

func (p *Part2Position) Forward(n int) {
	p.x += n
	p.z += p.aim * n
}

func (p *Part2Position) Result() int {
	return p.x * p.z
}
