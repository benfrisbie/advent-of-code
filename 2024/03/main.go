package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

// ReadInput reads the input file
func ReadInput(inputPath string) Data {
	r, err := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	if err != nil {
		panic(err)
	}

	var data Data
	// Open input file
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read file line by line
	enabled := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		matches := r.FindAllStringSubmatch(s, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
				continue
			}
			if match[0] == "don't()" {
				enabled = false
				continue
			}
			a, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			data.mul = append(data.mul, mul{x: a * b, enabled: enabled})
		}
	}
	return data
}

type Data struct {
	mul []mul
}
type mul struct {
	x       int
	enabled bool
}

func Part1(data Data) int {
	sum := 0
	for _, m := range data.mul {
		sum += m.x
	}
	return sum
}

func Part2(data Data) int {
	sum := 0
	for _, m := range data.mul {
		if m.enabled {
			sum += m.x
		}
	}
	return sum
}
