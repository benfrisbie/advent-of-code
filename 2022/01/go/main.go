package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Read input
	data, err := ReadInput("../input.txt")
	if err != nil {
		panic(err)
	}

	// Sort
	sort.Ints(data)

	// Part 1: Find the elf with the most calories
	fmt.Printf("Part 1 = %d\n", data[len(data)-1])

	// Part 2: Find the total of the three elves with the most calories
	sum := 0
	for _, i := range data[len(data)-3:] {
		sum += i
	}
	fmt.Printf("Part 2 = %d\n", sum)
}

// ReadInput reads the input file of groups of ints into a slice of int sums
func ReadInput(inputPath string) ([]int, error) {
	// Open input file
	var data []int
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Loop through input and sum each group of ints
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			data = append(data, sum)
			sum = 0
			continue
		}
		i, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		sum += i
	}
	return data, nil
}
