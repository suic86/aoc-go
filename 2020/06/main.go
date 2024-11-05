package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("input.data"))
	fmt.Printf("Part 2: %d\n", part2("input.data"))
}

func part1(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	union := map[rune]bool{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			total += len(union)
			union = map[rune]bool{}
		}
		for _, r := range t {
			union[r] = true
		}
	}
	return total + len(union)
}

func part2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	first := true
	intersection := map[rune]bool{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			total += len(intersection)
			first = true
			intersection = map[rune]bool{}
			continue
		}
		if first {
			for _, r := range t {
				intersection[r] = true
			}
			first = false
			continue
		}
		i := map[rune]bool{}
		for _, r := range t {
			if intersection[r] {
				i[r] = true
			}
		}
		intersection = i
	}

	return total + len(intersection)
}
