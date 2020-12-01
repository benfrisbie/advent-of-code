package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 2020

	// Read input
	data, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1: Find the pair that sums to 2020
	a, b, err := findPairThatSumsTo(sum, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pair = (%d, %d), product = %d\n", a, b, a*b)

	// Part 2: Find the triplet that sums to 2020
	a, b, c, err := findTripletThatSumsTo(sum, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("triplet = (%d, %d, %d), product = %d\n", a, b, c, a*b*c)
}

func readInput(inputPath string) ([]int, error) {
	// Read ints from a file with each on a newline
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

func findPairThatSumsTo(sum int, data []int) (int, int, error) {
	// Find the pair of numbers that adds up to the given sum
	found := make(map[int]bool)
	for i := 0; i < len(data); i++ {
		need := sum - data[i]
		if _, ok := found[need]; ok {
			return data[i], need, nil
		}
		found[data[i]] = true
	}
	return 0, 0, fmt.Errorf("No pair found")
}

func findTripletThatSumsTo(sum int, data []int) (int, int, int, error) {
	// Find the triplet of numbers that adds up to the given sum
	for i := 0; i < len(data); i++ {
		found := make(map[int]bool)
		for j := i; j < len(data); j++ {
			need := sum - data[i] - data[j]
			if _, ok := found[need]; ok {
				return data[i], data[j], sum - data[i] - data[j], nil
			}
			found[data[j]] = true
		}
	}
	return 0, 0, 0, fmt.Errorf("No triplet found")
}
