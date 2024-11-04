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
	{"example", 514579, "test.data"},
	{"input", 1020036, "input.data"},
}

var tests2 = []struct {
	name string
	want int
	file string
}{
	{"example", 241861950, "test.data"},
	{"input", 286977330, "input.data"},
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
