package main

import (
	"testing"

	"github.com/suic86/aoc-go/util"
)

var tests1 = []struct {
	name string
	want int
	file string
}{
	{"example", 2, "test.data"},
	{"input", 445, "input.data"},
}

func TestPart1(t *testing.T) {
	for _, test := range tests1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(util.ReadFile(test.file))
			if got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

var tests2 = []struct {
	name string
	want int
	file string
}{
	{"example", 1, "test.data"},
	{"input", 491, "input.data"},
}

func TestPart2(t *testing.T) {
	for _, test := range tests2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(util.ReadFile(test.file))
			if got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}

var lines1 = []struct {
	line  string
	valid bool
}{
	{"1-3 a: abcde", true},
	{"1-3 b: cdefg", false},
	{"2-9 c: ccccccccc", true},
}

func TestValidPassword1(t *testing.T) {
	for _, tt := range lines1 {
		t.Run(tt.line, func(*testing.T) {
			got := validPassword1(tt.line)
			if got != tt.valid {
				t.Errorf("got: %v, want: %v", got, tt.valid)
			}
		})
	}
}

var lines2 = []struct {
	line  string
	valid bool
}{
	{"1-3 a: abcde", true},
	{"1-3 b: cdefg", false},
	{"2-9 c: ccccccccc", false},
}

func TestValidPassword2(t *testing.T) {
	for _, tt := range lines2 {
		t.Run(tt.line, func(*testing.T) {
			got := validPassword2(tt.line)
			if got != tt.valid {
				t.Errorf("got: %v, want: %v", got, tt.valid)
			}
		})
	}
}
