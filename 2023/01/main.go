package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

// ReadInput reads the input file
func ReadInput(inputPath string) []string {
	var data []string
	// Open input file
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		data = append(data, s)
	}
	return data
}

var part1Digits []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Part1(data []string) int {
	sum := 0
	for _, s := range data {
		first, _ := FindFirstOccurrenceFromList(s, part1Digits)
		last, _ := FindLastOccurrenceFromList(s, part1Digits)
		num, _ := strconv.Atoi(first + last)
		sum += num
	}
	return sum
}

var part2Digits []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Part2(data []string) int {
	sum := 0
	for _, s := range data {
		first, _ := FindFirstOccurrenceFromList(s, part2Digits)
		first = ReplaceAlphaDigitWithNumericalDigit(first)
		last, _ := FindLastOccurrenceFromList(s, part2Digits)
		last = ReplaceAlphaDigitWithNumericalDigit(last)
		num, _ := strconv.Atoi(first + last)
		sum += num
	}
	return sum
}

func ReplaceAlphaDigitWithNumericalDigit(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return s
	}
}

func FindFirstOccurrenceFromList(s string, substrs []string) (string, int) {
	index := -1
	value := ""
	for _, d := range substrs {
		i := strings.Index(s, d)
		if i != -1 && (index == -1 || i < index) {
			index = i
			value = d
		}
	}
	return value, index
}

func FindLastOccurrenceFromList(s string, substrs []string) (string, int) {
	index := -1
	value := ""
	for _, d := range substrs {
		i := strings.LastIndex(s, d)
		if i != -1 && i > index {
			index = i
			value = d
		}
	}
	return value, index
}
