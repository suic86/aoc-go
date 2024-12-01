package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/suic86/aoc-go/util"
)

func main() {
	fmt.Printf("Part 1: %v\n", part1("input.data"))
	fmt.Printf("Part 2: %v\n", part2("input.data"))
}

func part1(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	left := []int{}
	right := []int{}
	for sc.Scan() {
		line := sc.Text()
		fs := strings.Split(line, "   ")
		l, err := strconv.Atoi(fs[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, l)
		r, err := strconv.Atoi(fs[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)
	total := 0
	for i := range left {
		total += util.IntDiffAbs(left[i], right[i])
	}
	return total
}

func part2(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	left := []int{}
	counter := map[int]int{}

	for sc.Scan() {
		line := sc.Text()
		fs := strings.Split(line, "   ")
		l, err := strconv.Atoi(fs[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, l)
		r, err := strconv.Atoi(fs[1])
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := counter[r]; !ok {
			counter[r] = 1
		} else {
			counter[r]++
		}
	}
	total := 0
	for _, l := range left {
		if val, ok := counter[l]; ok {
			total += l * val
		}
	}
	return total
}
