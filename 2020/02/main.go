package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"os"
	"path"
	"runtime"
)

func main() {
	fmt.Printf("Part 01: %v", part1(func() string {
		_, filename, _, ok := runtime.Caller(1)
		if !ok {
			panic("Could not find Caller of util.ReadFile")
		}
		absPath := path.Join(path.Dir(filename), "input.data")
		content, err := os.ReadFile(absPath)
		if err != nil {
			panic(err)
		}
		strContent := string(content)
		return strings.TrimRight(strContent, "\n")
	}()))
}

var lineParser = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func validPassword1(line string) bool {
	groups := lineParser.FindStringSubmatch(line)
	l := parseInt(groups[1])
	h := parseInt(groups[2])
	c := rune(groups[3][0])
	p := groups[4]

	ms := 0
	for _, n := range p {
		if n == c {
			ms++
		}
	}
	return l <= ms && ms <= h
}

func validPassword2(line string) bool {
	groups := lineParser.FindStringSubmatch(line)
	l := parseInt(groups[1])
	h := parseInt(groups[2])
	c := groups[3][0]
	p := groups[4]

	return (p[l-1] == c) != (p[h-1] == c)
}

func part1(data string) int {
	validPasswords := 0
	sc := bufio.NewScanner(strings.NewReader(data))
	for sc.Scan() {
		if validPassword1(sc.Text()) {
			validPasswords++
		}
	}
	return validPasswords
}

func part2(data string) int {
	validPasswords := 0
	sc := bufio.NewScanner(strings.NewReader(data))
	for sc.Scan() {
		if validPassword2(sc.Text()) {
			validPasswords++
		}
	}
	return validPasswords
}
