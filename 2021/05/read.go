package main

import (
	"bufio"
	"strings"
)

func getLines(s *bufio.Scanner) (lines []Line, maxX, maxY int) {
	lines = make([]Line, 0, 500)
	for s.Scan() {
		l := s.Text()
		parts := strings.Split(l, " -> ")

		startCoords, endCoords := strings.Split(parts[0], ","), strings.Split(parts[1], ",")
		line := Line{
			Start: Point{atoi(startCoords[0]), atoi(startCoords[1])},
			End:   Point{atoi(endCoords[0]), atoi(endCoords[1])},
		}

		if line.Start.X > maxX {
			maxX = line.Start.X
		}

		if line.End.X > maxX {
			maxX = line.End.X
		}

		if line.Start.Y > maxY {
			maxY = line.Start.Y
		}

		if line.End.Y > maxY {
			maxY = line.End.Y
		}

		lines = append(lines, line)
	}

	return lines, maxX, maxY
}
