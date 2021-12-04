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

			if hasWon(marked[x], board, pos) {
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
