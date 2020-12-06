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
	data, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1
	fmt.Println("Part 1 = ", IntCodeProgram(data, 12, 2))

	// Part 2
	fmt.Println("Part 2 = ", Part2(data, 19690720))
}

func IntCodeProgram(dataIn []int, noun int, verb int) int {
	data := make([]int, len(dataIn))
	copy(data, dataIn)
	data[1] = noun
	data[2] = verb
	for ip := 0; ; ip += 4 {
		opcode := data[ip]
		switch opcode {
		case 1:
			data[data[ip+3]] = data[data[ip+1]] + data[data[ip+2]]
		case 2:
			data[data[ip+3]] = data[data[ip+1]] * data[data[ip+2]]
		case 99:
			return data[0]
		default:
			continue
		}
	}
}

func Part2(data []int, desiredOutput int) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			output := IntCodeProgram(data, noun, verb)
			if output == desiredOutput {
				return 100*noun + verb
			}
		}
	}
	return 0
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
		for _, s := range strings.Split(scanner.Text(), ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			data = append(data, i)
		}
	}
	return data, nil
}
