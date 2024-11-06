package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %v\n", part1("input.data"))
	fmt.Printf("Part 2: %v\n", part2("input.data"))
}

var passportFields = []string{
	"byr", // (Birth Year)
	"iyr", // (Issue Year)
	"eyr", // (Expiration Year)
	"hgt", // (Height)
	"hcl", // (Hair Color)
	"ecl", // (Eye Color)
	"pid", // (Passport ID)
	"cid", // (Country ID)
}

func part1(path string) int {
	valid := 0
	for _, p := range parseData(path) {
		if hasAllFields(p) {
			valid++
		}
	}
	return valid
}

func part2(path string) int {
	valid := 0
	for _, p := range parseData(path) {
		if validPassport(p) {
			valid++
		}
	}
	return valid
}

func parsePassport(data []string) map[string]string {
	passport := map[string]string{}
	for _, line := range data {
		fields := strings.Split(line, " ")
		for _, field := range fields {
			f := strings.Split(field, ":")
			if len(f) != 2 {
				panic(fmt.Sprintf("invalid field: >%v<\n", field))
			}
			passport[f[0]] = f[1]
		}
	}
	return passport
}

func parseData(path string) []map[string]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	passports := []map[string]string{}
	pdata := []string{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		l := sc.Text()
		if l == "" {
			passports = append(passports, parsePassport(pdata))
			pdata = []string{}
			continue
		}
		pdata = append(pdata, l)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	if len(pdata) != 0 {
		passports = append(passports, parsePassport(pdata))
	}
	return passports
}

func validPassport(p map[string]string) bool {
	return hasAllFields(p) &&
		validByr(p["byr"]) &&
		validEcl(p["ecl"]) &&
		validEyr(p["eyr"]) &&
		validHcl(p["hcl"]) &&
		validHgt(p["hgt"]) &&
		validIyr(p["iyr"]) &&
		validPid(p["pid"])
}

func hasAllFields(p map[string]string) bool {
	for _, f := range passportFields {
		if _, ok := p[f]; !ok && f != "cid" {
			return false
		}
	}
	return true
}

func valueInRange(value string, lower int, upper int) bool {
	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return lower <= v && v <= upper
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func validByr(year string) bool {
	return valueInRange(year, 1920, 2002)
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func validIyr(year string) bool {
	return valueInRange(year, 2010, 2020)
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func validEyr(year string) bool {
	return valueInRange(year, 2020, 2030)
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func validHgt(height string) bool {
	if strings.HasSuffix(height, "cm") {
		return valueInRange(strings.TrimSuffix(height, "cm"), 150, 193)
	}
	if strings.HasSuffix(height, "in") {
		return valueInRange(strings.TrimSuffix(height, "in"), 59, 76)
	}
	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func validHcl(color string) bool {
	if len(color) != 7 || color[0] != '#' {
		return false
	}
	for _, r := range color[1:] {
		if !('0' <= r && r <= '9' || 'a' <= r && r <= 'f') {
			return false
		}
	}
	return true
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validEcl(color string) bool {
	switch color {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func validPid(color string) bool {
	if len(color) != 9 {
		return false
	}
	for _, r := range color {
		if !('0' <= r && r <= '9') {
			return false
		}
	}
	return true
}
