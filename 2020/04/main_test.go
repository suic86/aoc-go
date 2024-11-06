package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var tests1 = []struct {
	name string
	path string
	want int
}{
	{"example", "test.data", 2},
	{"input", "input.data", 242},
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

var tests2 = []struct {
	name string
	path string
	want int
}{
	{"input", "input.data", 186},
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

var testsParsePassport = []struct {
	name    string
	rawData []string
	want    map[string]string
}{
	{
		"valid",
		[]string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:183cm",
		},
		map[string]string{
			"ecl": "gry",
			"pid": "860033327",
			"eyr": "2020",
			"hcl": "#fffffd",
			"byr": "1937",
			"iyr": "2017",
			"cid": "147",
			"hgt": "183cm",
		},
	},
	{
		"multiline",
		[]string{
			"hcl:#ae17e1 iyr:2013",
			"eyr:2024",
			"ecl:brn pid:760753108 byr:1931",
			"hgt:179cm",
		},
		map[string]string{
			"hcl": "#ae17e1",
			"iyr": "2013",
			"eyr": "2024",
			"ecl": "brn",
			"pid": "760753108",
			"byr": "1931",
			"hgt": "179cm",
		},
	},
}

func TestParsePassport(t *testing.T) {
	for _, tt := range testsParsePassport {
		t.Run(tt.name, func(t *testing.T) {
			parsed := parsePassport(tt.rawData)
			if diff := cmp.Diff(tt.want, parsed); diff != "" {
				t.Errorf("parse missmatch (-want +got):\n%s", diff)
			}
		})
	}
}
