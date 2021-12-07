package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.ReadFile("input.txt")

	fmt.Println(Part1(string(f)))
	fmt.Println(Part2(string(f)))
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
