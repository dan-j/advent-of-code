package main

func Part1(input []Line) int {
	actualDefs := make(map[int]Set)
	answer := 0
	for _, line := range input {
		for _, s := range line.Digits {
			// digits 1, 4, 7 and 8 have a unique number of signals so we can deduce those straight away
			switch len(s) {
			case 2:
				actualDefs[1] = NewSet(s)
				answer++
			case 3:
				actualDefs[7] = NewSet(s)
				answer++
			case 4:
				actualDefs[4] = NewSet(s)
				answer++
			case 7:
				actualDefs[8] = NewSet(s)
				answer++
			}
		}
	}

	return answer
}
