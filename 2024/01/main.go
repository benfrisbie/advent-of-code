package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

// ReadInput reads the input file
func ReadInput(inputPath string) Data {
	data := Data{
		left:  make([]int, 0),
		right: make([]int, 0),
	}
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
		ss := strings.Split(s, "   ")
		i, err := strconv.Atoi(ss[0])
		if err != nil {
			panic(err)
		}
		data.left = append(data.left, i)
		i, err = strconv.Atoi(ss[1])
		if err != nil {
			panic(err)
		}
		data.right = append(data.right, i)
	}
	return data
}

type Data struct {
	left  []int
	right []int
}

func Part1(data Data) int {
	sum := 0
	slices.Sort(data.left)
	slices.Sort(data.right)
	for i := 0; i < len(data.left); i++ {
		diff := data.left[i] - data.right[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func Part2(data Data) int {
	sum := 0
	rightCounts := make(map[int]int)
	for _, r := range data.right {
		rightCounts[r]++
	}
	for _, l := range data.left {
		sum += l * rightCounts[l]
	}

	return sum
}
