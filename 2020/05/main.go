package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	// Read input
	data, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Sort the seat ids
	sort.Ints(data)

	// Part 1: max seat id is at end of slice since its sorted
	fmt.Println("Part 1: max seat id = ", data[len(data)-1])

	// Part 2
	fmt.Println("Part 2: my seat id = ", FindMySeatID(data))
}

// FindMySeatID finds my seat id by looking for the missing id in the middle of the sorted slice
func FindMySeatID(data []int) int {
	lastSeatID := data[0]
	for _, seatID := range data[1:] {
		if seatID != lastSeatID+1 {
			return seatID - 1
		}
		lastSeatID = seatID
	}

	return 0
}

// ReadInput reads the input file into a slice of ints representing boarding pass seat ids
func ReadInput(inputPath string) ([]int, error) {
	// Loop through file and convert each line to a string
	var data []int
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bp := scanner.Text()
		row := 0
		for i, s := range bp[:7] {
			if s == 'B' {
				row += int(math.Pow(2, float64(6-i)))
			}
		}
		col := 0
		for i, s := range bp[7:] {
			if s == 'R' {
				col += int(math.Pow(2, float64(2-i)))
			}
		}
		data = append(data, row*8+col)
	}
	return data, nil
}
