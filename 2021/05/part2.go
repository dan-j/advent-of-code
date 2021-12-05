package main

import (
	"bufio"
	"bytes"
	"os"
)

func Part2() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	lines, sizeX, sizeY := getLines(s)

	// lines = filterStraight(lines)

	grid := make([][]int, sizeX+1)
	for i := range grid {
		grid[i] = make([]int, sizeY+1)
	}

	for _, line := range lines {
		if line.IsVertical() {
			x := line.Start.X

			var minY, maxY int
			if line.Start.Y < line.End.Y {
				minY = line.Start.Y
				maxY = line.End.Y
			} else {
				minY = line.End.Y
				maxY = line.Start.Y
			}

			for y := minY; y <= maxY; y++ {
				grid[x][y]++
			}
		} else if line.IsHorizontal() {
			y := line.Start.Y

			var minX, maxX int
			if line.Start.X <= line.End.X {
				minX = line.Start.X
				maxX = line.End.X
			} else {
				minX = line.End.X
				maxX = line.Start.X
			}

			for x := minX; x <= maxX; x++ {
				grid[x][y]++
			}
		} else {
			var xFn func(x int) int
			if line.Start.X <= line.End.X {
				xFn = func(x int) int { return x + 1 }
			} else {
				xFn = func(x int) int { return x - 1 }
			}

			var minY, maxY int
			var yFn func(x int) int
			if line.Start.Y < line.End.Y {
				yFn = func(y int) int { return y + 1 }
				minY = line.Start.Y
				maxY = line.End.Y
			} else {
				yFn = func(y int) int { return y - 1 }
				minY = line.End.Y
				maxY = line.Start.Y
			}

			x, y := line.Start.X, line.Start.Y
			for i := 0; i <= maxY-minY; i++ {
				grid[x][y]++
				x = xFn(x)
				y = yFn(y)
			}
		}
	}

	var result int
	for i := range grid {
		for j := range grid[i] {
			// fmt.Printf("%d", grid[j][i])
			if grid[i][j] >= 2 {
				result++
			}
		}
		// fmt.Println()
	}

	return result
}
