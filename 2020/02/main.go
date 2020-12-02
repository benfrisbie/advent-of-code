package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input
	data, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	fmt.Println("Part 1: # of valid passwords = ", CountValidPasswords(data, ValidatePasswordPart1))

	// Part 2
	fmt.Println("Part 2: # of valid passwords = ", CountValidPasswords(data, ValidatePasswordPart2))
}

// PasswordValidator contains the info from the input file needed to check if a password is valid
type PasswordValidator struct {
	Num1     int
	Num2     int
	Letter   string
	Password string
}

// ValidatePassword defines a function type for validating passwords
type ValidatePassword func(pv PasswordValidator) bool

// ValidatePasswordPart1 validates a password for part 1 of the puzzle
func ValidatePasswordPart1(pv PasswordValidator) bool {
	count := strings.Count(pv.Password, pv.Letter)
	if count < pv.Num1 || count > pv.Num2 {
		return false
	}
	return true
}

// ValidatePasswordPart2 validates a password for part 2 of the puzzle
func ValidatePasswordPart2(pv PasswordValidator) bool {
	p1 := string(pv.Password[pv.Num1-1]) == pv.Letter
	p2 := string(pv.Password[pv.Num2-1]) == pv.Letter

	if (p1 || p2) && !(p1 && p2) {
		return true
	}
	return false
}

// CountValidPasswords counts the number of valid passwords in the data slice
func CountValidPasswords(data []PasswordValidator, fn ValidatePassword) int {
	valid := 0
	for _, pv := range data {
		if fn(pv) {
			valid++
		}
	}
	return valid
}

// ReadInput reads the input file into a slice of PasswordValidator structs
func ReadInput(inputPath string) ([]PasswordValidator, error) {
	// Compile regex
	r, err := regexp.Compile("(\\d+)-(\\d+)\\s+(.)\\:\\s+(.*)")
	if err != nil {
		return nil, err
	}

	// Loop through the file and generate PasswordValidator structs
	var data []PasswordValidator
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, err
		}
		data = append(data, PasswordValidator{Num1: num1, Num2: num2, Letter: match[3], Password: match[4]})
	}
	return data, nil
}
