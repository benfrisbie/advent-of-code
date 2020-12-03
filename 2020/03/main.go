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

	// Solve Part 1
	fmt.Println("Part1: Trees encountered = ", CountTreesOnSlope(data, 3, 1))

	// Solve Part 2
	product := CountTreesOnSlope(data, 1, 1)
	product *= CountTreesOnSlope(data, 3, 1)
	product *= CountTreesOnSlope(data, 5, 1)
	product *= CountTreesOnSlope(data, 7, 1)
	product *= CountTreesOnSlope(data, 1, 2)
	fmt.Println("Part 2: Trees encountered = ", product)
}

// CountTreesOnSlope counts the number of trees encountered on a slope. The slope is to move down and right the specified numbers
func CountTreesOnSlope(data [][]bool, right int, down int) int {
	rowLen := len(data[0])
	count := 0
	x := right
	for y := down; y < len(data); y += down {
		if data[y][x%rowLen] {
			count++
		}
		x += right
	}
	return count
}

// ReadInput reads the input file into a 2d slice of bool
func ReadInput(inputPath string) ([][]bool, error) {
	var data [][]bool
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []bool
		for _, s := range scanner.Text() {
			row = append(row, s == '#')
		}
		data = append(data, row)
	}
	return data, nil
}
