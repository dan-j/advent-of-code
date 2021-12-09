package main

func Part1(input [][]int) int {
	var risk int

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

			risk += v + 1
		}
	}

	return risk
}
