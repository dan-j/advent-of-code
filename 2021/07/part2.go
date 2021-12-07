package main

import (
	"math"
	"sort"
	"strings"
)

func Part2(input string) int {
	split := strings.Split(input, ",")
	positions := make([]int, len(split))
	for i, p := range split {
		positions[i] = atoi(p)
	}

	sort.Ints(positions)

	maxPos := positions[len(positions)-1]
	optimumFuel := sum(maxPos) * len(positions)
	for i := range make([]struct{}, maxPos) {

		var fuel int
		for _, p := range positions {
			n := int(math.Abs(float64(p - i)))
			fuel += sum(n)
		}

		if fuel < optimumFuel {
			optimumFuel = fuel
		} else {
			break
		}
	}

	return optimumFuel
}

func sum(n int) int {
	return (n * (n + 1)) / 2
}
