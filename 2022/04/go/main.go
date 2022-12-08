package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input
	data, err := ReadInput("../input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	fmt.Println("Part1 = ", Part1(data))

	// Part 2
	fmt.Println("Part2 = ", Part2(data))
}

func Part1(pairs []Pair) int {
	overlappingPairs := 0
	for _, pair := range pairs {
		if pair.First.MinAssignment >= pair.Second.MinAssignment && pair.First.MaxAssignment <= pair.Second.MaxAssignment {
			overlappingPairs += 1
		} else if pair.Second.MinAssignment >= pair.First.MinAssignment && pair.Second.MaxAssignment <= pair.First.MaxAssignment {
			overlappingPairs += 1
		}
	}
	return overlappingPairs
}

func Part2(pairs []Pair) int {
	overlappingPairs := 0
	for _, pair := range pairs {
		if pair.First.MinAssignment <= pair.Second.MaxAssignment && pair.First.MaxAssignment >= pair.Second.MinAssignment {
			overlappingPairs += 1
		} else if pair.Second.MinAssignment <= pair.First.MaxAssignment && pair.Second.MaxAssignment >= pair.First.MinAssignment {
			overlappingPairs += 1
		}
	}
	return overlappingPairs
}

type Elf struct {
	MinAssignment, MaxAssignment int
}
type Pair struct {
	First, Second Elf
}

// ReadInput reads the input file
func ReadInput(inputPath string) ([]Pair, error) {
	var data []Pair
	// Open input file
	file, err := os.Open(inputPath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	var p Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		s0 := strings.Split(s[0], "-")
		s1 := strings.Split(s[1], "-")
		p.First.MinAssignment, _ = strconv.Atoi(s0[0])
		p.First.MaxAssignment, _ = strconv.Atoi(s0[1])
		p.Second.MinAssignment, _ = strconv.Atoi(s1[0])
		p.Second.MaxAssignment, _ = strconv.Atoi(s1[1])
		data = append(data, p)
	}
	return data, nil
}
