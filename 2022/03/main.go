package main

import (
	"bufio"
	"fmt"
	"os"
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

func Part1(rucksacks []Rucksack) int {
	totalPriority := 0
	for _, rucksack := range rucksacks {
		for _, r := range rucksack.FirstCompartment {
			if strings.ContainsRune(rucksack.SecondCompartment, r) {
				totalPriority += getRunePriority(r)
				break
			}
		}
	}
	return totalPriority
}

func Part2(rucksacks []Rucksack) int {
	totalPriority := 0
	for i := 0; i < len(rucksacks); i += 3 {
		for _, r := range rucksacks[i].FirstCompartment + rucksacks[i].SecondCompartment {
			if !strings.ContainsRune(rucksacks[i+1].FirstCompartment+rucksacks[i+1].SecondCompartment, r) {
				continue
			}
			if !strings.ContainsRune(rucksacks[i+2].FirstCompartment+rucksacks[i+2].SecondCompartment, r) {
				continue
			}
			totalPriority += getRunePriority(r)
			break
		}
	}
	return totalPriority
}

func getRunePriority(r rune) int {
	if r >= 97 && r <= 122 { // Rune 'a' - 'z' are 97 - 122
		return int(r) - 96 // subtract 96 to get priority 1 - 26
	} else { // Rune 'A' - 'Z' are 65 - 90
		return int(r) - 64 + 26 // subtract 64 to get priority 27 - 52
	}
}

type Rucksack struct {
	FirstCompartment, SecondCompartment string
}

// ReadInput reads the input file
func ReadInput(inputPath string) ([]Rucksack, error) {
	var data []Rucksack
	// Open input file
	file, err := os.Open(inputPath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	var r Rucksack
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		r.FirstCompartment = text[:len(text)/2]
		r.SecondCompartment = text[len(text)/2:]
		data = append(data, r)
	}
	return data, nil
}
