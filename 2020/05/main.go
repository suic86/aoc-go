package main

import (
	"fmt"

	"github.com/suic86/aoc-go/util"
)

func main() {
	fmt.Printf("Solution 01: %d\n", part1("input.data"))
	fmt.Printf("Solution 02: %d\n", part2("input.data"))
}

func convertToSeat(code string) int {
	seat := 0
	for i, r := range code {
		if r == 'B' {
			// Convert the first seven chars to int
			// 1 << (7 - i - 1)
			seat += 8 * (1 << (6 - i))
		} else if r == 'R' {
			// Convert the 8th to 10th chars to int
			// 1 << (3 - i - 1 + 6)
			seat += 1 << (9 - i)
		}
	}
	return seat
}

func part1(path string) int {
	max := 0
	for _, code := range util.ReadFileToStringArray(path) {
		seat := convertToSeat(code)
		if seat > max {
			max = seat
		}
	}
	return max
}

func part2(path string) int {
	// Explanation: https://stackoverflow.com/a/18335854
	acc := 0
	for _, c := range util.ReadFileToStringArray(path) {
		acc = acc ^ convertToSeat(c)
	}
	return acc
}
