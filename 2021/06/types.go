package main

import (
	"strconv"
	"strings"
)

type Fishes []Fish

func (f Fishes) String() string {
	var sb strings.Builder
	for i, fish := range f {
		sb.WriteString(fish.String())
		if i < len(f)-1 {
			sb.WriteRune(',')
		}
	}
	return sb.String()
}

type Fish struct {
	Timer int
}

func (f *Fish) Decrement() (offspring Fish, spawned bool) {
	if f.Timer == 0 {
		f.Timer = 6
		return Fish{Timer: 8}, true
	}

	f.Timer--
	return offspring, false
}

func (f *Fish) String() string {
	return strconv.Itoa(f.Timer)
}
