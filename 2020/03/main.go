package main

import (
	"fmt"

	"github.com/suic86/aoc-go/util"
)

func main() {
	fmt.Printf("Part 1: %v\n", part1("input.data"))
	fmt.Printf("Part 2: %v\n", part2("input.data"))
}

func part1(path string) int {
	return countTrees(util.ReadFileToStringArray(path), 3, 1)
}

func part2(path string) int {
	var traverses = []struct {
		rightStep int
		downStep  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	data := util.ReadFileToStringArray(path)
	total := 1
	for _, t := range traverses {
		total *= countTrees(data, t.rightStep, t.downStep)
	}
	return total
}

func countTrees(data []string, rightStep int, downStep int) int {
	h := len(data)
	w := len(data[0])
	x := 0
	trees := 0
	for y := 0; y < h; y += downStep {
		if data[y][x] == '#' {
			trees++
		}
		x = (x + rightStep) % w
	}
	return trees
}
