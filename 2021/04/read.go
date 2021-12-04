package main

import (
	"bufio"
	"strings"
)

func readInput(s *bufio.Scanner) (nums []string, boards [][][]string, indexes []map[string][]int) {
	s.Scan()
	nums = strings.Split(s.Text(), ",")
	for s.Scan() {
		board, index := readBoard(s)
		boards = append(boards, board)
		indexes = append(indexes, index)
	}
	return nums, boards, indexes
}

func readBoard(s *bufio.Scanner) (board [][]string, index map[string][]int) {
	board = make([][]string, 5)
	index = make(map[string][]int)
	for i := 0; i < 5; i++ {
		s.Scan()
		line := s.Text()

		if line == "" {
			return board, index
		}

		row := strings.Split(line, " ")

		board[i] = make([]string, 5)

		// single-digit numbers have padding so end up with additional empty items in row, remove them.
		for j, n := range row {
			if n == "" {
				copy(row[j:], row[j+1:])
				row = row[:len(row)-1]
			}
		}

		for j, n := range row {
			board[i][j] = n
			index[n] = []int{i, j}
		}
	}

	return board, index
}
