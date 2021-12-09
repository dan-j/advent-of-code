package main

import "sort"

func Part2(input [][]int) int {
	var basinCentres [][]int

	for i := range input {
		for j := range input[i] {
			v := input[i][j]

			if i-1 >= 0 && v >= input[i-1][j] {
				continue
			}

			if i+1 < len(input) && v >= input[i+1][j] {
				continue
			}

			if j-1 >= 0 && v >= input[i][j-1] {
				continue
			}

			if j+1 < len(input[i]) && v >= input[i][j+1] {
				continue
			}

			basinCentres = append(basinCentres, []int{i, j})
		}
	}

	basinSizes := make([]int, len(basinCentres))
	for n := range basinCentres {
		i, j := basinCentres[n][0], basinCentres[n][1]
		processed := map[Point]Process{}
		sum := basinSize(input, i, j, processed)
		basinSizes[n] = sum
	}

	sort.Ints(basinSizes)
	result := 1
	for _, size := range basinSizes[len(basinSizes)-3:] {
		result *= size
	}

	return result
}

// basinSize traverses the basin using recursion. Travelling firstly in a horizontal direction, and for each element
// traverses vertically, and then for each element traverses horizontally etc. A map is kept up to date whether
// the element being visited has already been traversed horizontally or vertically, allowing the recursion to stop.
func basinSize(input [][]int, i, j int, processed map[Point]Process) (sum int) {
	processed[Point{i, j}] = Both
	return 1 + basinWidth(input, i, j, processed) + basinLength(input, i, j, processed)
}

func basinWidth(input [][]int, i, j int, processed map[Point]Process) (size int) {
	if j > 0 {
		// go left
		for x := 1; j-x >= 0 && input[i][j-x] != 9; x++ {
			// this point hasn't been processed, increment size.
			if _, ok := processed[Point{i, j - x}]; !ok {
				// mark as being processed horizontally in case this same element is visited while traversing vertically
				processed[Point{i, j - x}] |= Horizontally
				size++
			}

			// if this element hasn't been traversed vertically, step into the vertical directory
			if processed[Point{i, j - x}]&Vertically != Vertically {
				processed[Point{i, j - x}] |= Vertically
				size += basinLength(input, i, j-x, processed)
			}
		}
	}

	if j < len(input[i]) {
		// go right
		for x := 1; j+x < len(input[i]) && input[i][j+x] != 9; x++ {
			if _, ok := processed[Point{i, j + x}]; !ok {
				size++
				processed[Point{i, j + x}] |= Horizontally
			}

			if processed[Point{i, j + x}]&Vertically != Vertically {
				processed[Point{i, j + x}] |= Vertically
				size += basinLength(input, i, j+x, processed)
			}
		}
	}

	return size
}

func basinLength(input [][]int, i, j int, processed map[Point]Process) (size int) {
	if i > 0 {
		// go up
		for x := 1; i-x >= 0 && input[i-x][j] != 9; x++ {
			if _, ok := processed[Point{i - x, j}]; !ok {
				size++
				processed[Point{i - x, j}] |= Vertically
			}

			if processed[Point{i - x, j}]&Horizontally != Horizontally {
				processed[Point{i - x, j}] |= Horizontally
				size += basinWidth(input, i-x, j, processed)
			}
		}
	}

	if i < len(input) {
		// go down
		for x := 1; i+x < len(input) && input[i+x][j] != 9; x++ {
			if _, ok := processed[Point{i + x, j}]; !ok {
				size++
				processed[Point{i + x, j}] |= Vertically
			}

			if processed[Point{i + x, j}]&Horizontally != Horizontally {
				processed[Point{i + x, j}] |= Horizontally
				size += basinWidth(input, i+x, j, processed)
			}
		}
	}

	return size
}

type Process int

const (
	None Process = iota << 1
	Vertically
	Horizontally

	Both = Vertically | Horizontally
)

type Point struct {
	i, j int
}
