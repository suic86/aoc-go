package main

import "testing"

var tests1 = []struct{
	name string
	want int
	file string
}{
	{"example", 7, "test.data"},
	{"input", 151, "input.data"},
}

func TestPart1(t *testing.T) {
    for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.file)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

var tests2 = []struct{
	name string
	want int
	file string
}{
	{"example", 336, "test.data"},
	{"input", 7540141059, "input.data"},
}

func TestPart2(t *testing.T) {
    for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.file)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
