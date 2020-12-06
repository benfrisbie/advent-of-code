package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Read input
	data, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	fmt.Println("Part 1: checksum = ", Part1(data))

	// Part 2
	fmt.Println("Part 2: common box id letters = ", Part2(data))
}

func Part1(ids []string) int {
	twoLettersCount := 0
	threeLettersCount := 0

	for _, id := range ids {
		letterCounts := make(map[rune]int)
		for _, r := range id {
			letterCounts[r]++
		}

		twoLetters := false
		threeLetters := false
		for _, v := range letterCounts {
			if v == 2 {
				twoLetters = true
			} else if v == 3 {
				threeLetters = true
			}
			if twoLetters && threeLetters {
				break
			}
		}
		if twoLetters {
			twoLettersCount++
		}
		if threeLetters {
			threeLettersCount++
		}
	}
	return twoLettersCount * threeLettersCount
}

func Part2(ids []string) string {
	sort.Strings(ids)
	for i := 0; i < len(ids)-2; i++ {
		diffIndex := -1
		for j := range ids[i] {
			if ids[i][j] != ids[i+1][j] {
				if diffIndex != -1 {
					diffIndex = -1
					break
				}
				diffIndex = j
			}
		}
		if diffIndex > -1 {
			return ids[i][:diffIndex] + ids[i][diffIndex+1:]
		}
	}
	return ""
}

// ReadInput reads the input file into a slice of strings
func ReadInput(inputPath string) ([]string, error) {
	var data []string

	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}
