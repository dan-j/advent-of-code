package main

import (
	"fmt"
	"strconv"
)

type Point struct {
	X, Y int
}

type Line struct {
	Start, End Point
}

func (l *Line) IsHorizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l *Line) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l *Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
