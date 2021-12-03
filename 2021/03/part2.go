package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func Part2() int64 {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))
	var oxygenLines, scrubberLines []string
	for s.Scan() {
		oxygenLines = append(oxygenLines, s.Text())
		scrubberLines = append(scrubberLines, s.Text())
	}

	for pos := 0; pos < 12; pos++ {
		oxygenLines = reduceLines(oxygenLines, pos, true)
		scrubberLines = reduceLines(scrubberLines, pos, false)
	}
	oxygen, _ := strconv.ParseInt(oxygenLines[0], 2, 64)
	scrubber, _ := strconv.ParseInt(scrubberLines[0], 2, 64)
	return oxygen * scrubber
}

func reduceLines(arr []string, pos int, majority bool) []string {
	if len(arr) == 1 {
		return arr
	}

	var ones, zeros []string
	for _, line := range arr {
		if line[pos] == '1' {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}

	if majority {
		if len(ones) >= len(zeros) {
			return ones
		}
		return zeros
	} else if len(ones) >= len(zeros) {
		return zeros
	}

	return ones
}
