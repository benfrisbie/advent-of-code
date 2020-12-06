package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input
	data, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1 = ", Part1(data))

	fmt.Println("Part 2 = ", Part2(data))
}

func Part1(data [][][]rune) int {
	sum := 0
	for _, group := range data {
		questionsAnswered := make(map[rune]bool)
		for _, individual := range group {
			for _, question := range individual {
				questionsAnswered[question] = true
			}
		}
		sum += len(questionsAnswered)
	}
	return sum
}

func Part2(data [][][]rune) int {
	sum := 0
	for _, group := range data {
		questionsAnswered := make(map[rune]int)
		for _, individual := range group {
			for _, question := range individual {
				questionsAnswered[question]++
			}
		}
		for _, count := range questionsAnswered {
			if count == len(group) {
				sum++
			}
		}
	}
	return sum
}

// ReadInput reads the input file
func ReadInput(inputPath string) ([][][]rune, error) {
	var data [][][]rune
	var group [][]rune
	var individual []rune

	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			data = append(data, group)
			group = make([][]rune, 0)
			continue
		}
		individual = make([]rune, 0)
		for _, r := range text {
			individual = append(individual, r)
		}
		group = append(group, individual)
	}
	data = append(data, group)
	return data, nil
}
