package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func Part1() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	nums, boards, indexes := readInput(s)

	marked := make([]map[string]struct{}, len(boards))
	for i := range marked {
		marked[i] = make(map[string]struct{})
	}

	for _, n := range nums {
		for x, board := range boards {
			marked[x][n] = struct{}{}
			pos, ok := indexes[x][n]

			if !ok {
				// number not in board
				continue
			}

			// check row
			won := true
			for i := 0; i < 5; i++ {
				if _, ok := marked[x][board[i][pos[1]]]; !ok {
					won = false
					break
				}
			}

			// check column
			won = true
			for i := 0; i < 5; i++ {
				if _, ok := marked[x][board[pos[0]][i]]; !ok {
					won = false
					break
				}
			}

			if won {
				var unmarkedSum int
				for i, row := range board {
					for j := range row {
						vs := board[i][j]
						if _, ok := marked[x][vs]; !ok {
							v, _ := strconv.Atoi(vs)
							unmarkedSum += v
						}
					}
				}

				nn, _ := strconv.Atoi(n)
				return unmarkedSum * nn
			}
		}
	}

	return 0
}
