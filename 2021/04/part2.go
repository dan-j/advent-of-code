package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func Part2() int {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	nums, boards, indexes := readInput(s)

	marked := make([]map[string]struct{}, len(boards))
	for i := range marked {
		marked[i] = make(map[string]struct{})
	}

	boardsWon := make(map[int]struct{})

	var lastNumber, lastBoardScore, lastBoard int
	for _, n := range nums {
		for x, board := range boards {
			if _, ok := boardsWon[x]; ok {
				// board's already won
				continue
			}

			pos, ok := indexes[x][n]

			if !ok {
				// number not in board
				continue
			}

			marked[x][n] = struct{}{}

			if hasWon(marked[x], board, pos) {
				boardsWon[x] = struct{}{}
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
				lastNumber = nn
				lastBoard = x
				lastBoardScore = unmarkedSum * nn
			}
		}
	}

	fmt.Println(lastNumber * lastBoard)
	return lastBoardScore
}
