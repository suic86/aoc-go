package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/suic86/aoc-go/util"
)

func main() {
	input := util.ReadFile("input.data")
	fmt.Printf("Solution 01: %v\n", part1(input))
	fmt.Printf("Solution 02: %v\n", part2(input))
}

func part1(input string) int {
	seen := map[int]bool{}
	ns := parseInput(input)
	for _, n := range ns {
		if seen[n] {
			return n * (2020 - n)
		}
		seen[2020-n] = true
	}
	return -1 // unreachable
}

func part2(input string) int {
	ns := parseInput(input)
	ln := len(ns)
	for i := range ln {
		for j := i + 1; j < ln; j++ {
			for k := range ln {
				if ns[i]+ns[j]+ns[k] == 2020 {
					return ns[i] * ns[j] * ns[k]
				}
			}
		}
	}

	return -1 // unreachable
}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	nums := []int{}
	for _, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}
