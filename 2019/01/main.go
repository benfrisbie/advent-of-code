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
	fmt.Println("Part 1 = ", Part1(data))

	// Part 2
	fmt.Println("Part 2 = ", Part2(data))
}

func Part1(data []int) int {
	sum := 0
	for _, mass := range data {
		sum += CalculateFuelForModule(mass)
	}
	return sum
}

func Part2(data []int) int {
	sum := 0
	for _, mass := range data {
		sum += CalculateFuelForModuleRecursive(mass)
	}
	return sum
}

func CalculateFuelForModuleRecursive(mass int) int {
	fuel := CalculateFuelForModule(mass)
	if fuel > 0 {
		return fuel + CalculateFuelForModuleRecursive(fuel)
	}
	return 0
}

func CalculateFuelForModule(mass int) int {
	return (mass / 3) - 2
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
