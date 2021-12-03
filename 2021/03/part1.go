package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

func Part1() int64 {
	f, _ := os.ReadFile("input.txt")
	s := bufio.NewScanner(bytes.NewReader(f))

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	var gamma, epsilon strings.Builder
	for pos := 0; pos < 12; pos++ {
		var one, zero int
		for _, line := range lines {
			if line[pos] == '1' {
				one++
			} else {
				zero++
			}
		}

		if one > zero {
			gamma.WriteRune('1')
			epsilon.WriteRune('0')
		} else {
			gamma.WriteRune('0')
			epsilon.WriteRune('1')
		}
	}

	gammaNum, _ := strconv.ParseInt(gamma.String(), 2, 64)
	epsilonNum, _ := strconv.ParseInt(epsilon.String(), 2, 64)

	return gammaNum * epsilonNum
}
