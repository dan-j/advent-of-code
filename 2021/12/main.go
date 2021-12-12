package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	fmt.Println(Part1(string(f)))
	fmt.Println(Part2(string(f)))
}
