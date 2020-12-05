package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Read input
	data, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Solve Part 1
	fmt.Println("Part 1: # of valid passports", CountValidPassports(data, ValidatePassportPart1))

	// Solve Part 2
	fmt.Println("Part 2: # of valid passports", CountValidPassports(data, ValidatePassportPart2))
}

// RequiredKeys are the keys required for a valid passport
var RequiredKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

// ValidatePassport defines a function type for passport validation
type ValidatePassport func(p map[string]string) bool

// ValidatePassportPart1 validates a passport for part 1
func ValidatePassportPart1(p map[string]string) bool {
	for _, neededKey := range RequiredKeys {
		if _, ok := p[neededKey]; !ok {
			return false
		}
	}
	return true
}

// ValidatePassportPart2 validates a passport for part 2
func ValidatePassportPart2(p map[string]string) bool {
	for _, neededKey := range RequiredKeys {
		if _, ok := p[neededKey]; !ok || !ValidateFieldPart2(neededKey, p[neededKey]) {
			return false
		}
	}
	return true
}

// ValidateFieldPart2 validates a given field for part 2
func ValidateFieldPart2(k string, v string) bool {
	switch k {
	case "byr":
		i, err := strconv.Atoi(v)
		return len(v) == 4 && err == nil && i >= 1920 && i <= 2002
	case "iyr":
		i, err := strconv.Atoi(v)
		return len(v) == 4 && err == nil && i >= 2010 && i <= 2020
	case "eyr":
		i, err := strconv.Atoi(v)
		return len(v) == 4 && err == nil && i >= 2020 && i <= 2030
	case "hgt":
		height, err := strconv.Atoi(v[:len(v)-2])
		if err != nil {
			return false
		}
		unit := v[len(v)-2:]
		if unit == "cm" {
			return height >= 150 && height <= 193
		} else if unit == "in" {
			return height >= 59 && height <= 76
		}
		return false
	case "hcl":
		_, err := hex.DecodeString(v[1:])
		return len(v) == 7 && err == nil && v[0] == '#'
	case "ecl":
		return v == "amb" || v == "blu" || v == "brn" || v == "gry" || v == "grn" || v == "hzl" || v == "oth"
	case "pid":
		_, err := strconv.Atoi(v)
		return len(v) == 9 && err == nil
	default:
		return true
	}
}

// CountValidPassports counts the number of valid passports given some ValidatePassport function
func CountValidPassports(passports []map[string]string, fn ValidatePassport) int {
	count := 0
	for _, p := range passports {
		if fn(p) {
			count++
		}
	}
	return count
}

// ReadInput reads the input file into a slice of map[string]string
func ReadInput(inputPath string) ([]map[string]string, error) {
	var passports []map[string]string
	passport := make(map[string]string)

	// Compile regex
	r, err := regexp.Compile("(\\w{3}):([\\w#]+)")
	if err != nil {
		return nil, err
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
			continue
		}

		for _, match := range r.FindAllStringSubmatch(text, -1) {
			passport[match[1]] = match[2]
		}
	}
	passports = append(passports, passport)

	return passports, nil
}
