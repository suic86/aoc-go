package main

import (
	"testing"
)

const inputData = "input.data"
const exampleData = "test.data"

func TestPart1(t *testing.T) {
	testCases := []struct {
		name string
		path string
		want int
	}{
		{"example", exampleData, 820},
		{"input", inputData, 987},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.path)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name string
		path string
		want int
	}{
		{"input", "input.data", 603},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.path)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestConvertToSeat(t *testing.T) {
	testCases := []struct {
		code string
		want int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
	}
	for _, tC := range testCases {
		t.Run(tC.code, func(t *testing.T) {
			got := convertToSeat(tC.code)
			if got != tC.want {
				t.Errorf("got: %v, want: %v", got, tC.want)
			}
		})
	}
}
