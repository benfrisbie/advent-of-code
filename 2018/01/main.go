package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input
	data := ReadInput("input.txt")

	// Part 1
	fmt.Println("Part 1: final frequency = ", Part1(data))

	// Part 2
	fmt.Println("Part 2: first repeating frequency = ", Part2(data))
}

// CalculateFinalFrequency calculates the final frequency after all the deltas are applied
func Part1(data []int) int {
	freq := 0
	for _, delta := range data {
		freq += delta
	}
	return freq
}

// Part2 finds the first frequency that appears twice
func Part2(data []int) int {
	found := make(map[int]bool)
	freq := 0
	for i := 0; ; i++ {
		freq += data[i%len(data)]
		if _, ok := found[freq]; ok {
			return freq
		}
		found[freq] = true
	}
}

// ReadInput reads the input file into a slice of ints
func ReadInput(inputPath string) []int {
	// Loop through file and convert each line to an int
	var data []int
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		data = append(data, i)
	}
	return data
}
