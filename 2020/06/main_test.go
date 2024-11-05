package main

import "testing"

var tests1 = []struct{
	name string
	path string
	want int
}{
	{"example", "test.data", 11},
	{"input", "input.data", 6630},
}

func TestPart1(t *testing.T) {
    for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.path)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

var tests2 = []struct{
	name string
	path string
	want int
}{
	{"example", "test.data", 6},
	{"input", "input.data", 3437},
}

func TestPart2(t *testing.T) {
    for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.path)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
