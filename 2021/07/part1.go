package main

import (
	"math"
	"sort"
	"strings"
)

func Part1(input string) int {
	split := strings.Split(input, ",")
	positions := make([]int, len(split))
	for i, p := range split {
		positions[i] = atoi(p)
	}

	sort.Ints(positions)

	optimumFuel := positions[len(positions)-1] * len(positions)
	for i := range make([]struct{}, positions[len(positions)-1]) {

		var fuel int
		for _, p := range positions {
			fuel += int(math.Abs(float64(p - i)))
		}

		if fuel < optimumFuel {
			optimumFuel = fuel
		} else {
			break
		}
	}

	return optimumFuel
}
