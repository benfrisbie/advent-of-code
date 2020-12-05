package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input
	data, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	fmt.Println("Part 1: final frequency = ", CalculateFinalFrequency(data))

	// Part 2
	fmt.Println("Part 2: first repeating frequency = ", FindFirstRepeatingFrequency(data))
}

// CalculateFinalFrequency calculates the final frequency after all the deltas are applied
func CalculateFinalFrequency(data []int) int {
	freq := 0
	for _, delta := range data {
		freq += delta
	}
	return freq
}

// FindFirstRepeatingFrequency finds the first frequency that appears twice
func FindFirstRepeatingFrequency(data []int) int {
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
func ReadInput(inputPath string) ([]int, error) {
	// Loop through file and convert each line to an int
	var data []int
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	return data, nil
}
