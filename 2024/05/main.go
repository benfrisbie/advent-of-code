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
		rules: make(map[int][]int),
	}
	// Open input file
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	section1 := true
	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			section1 = false
			continue
		}
		if section1 {
			ss := strings.Split(s, "|")
			i, err := strconv.Atoi(ss[0])
			if err != nil {
				panic(err)
			}
			j, err := strconv.Atoi(ss[1])
			if err != nil {
				panic(err)
			}
			data.rules[i] = append(data.rules[i], j)
		} else {
			var u []int
			ss := strings.Split(s, ",")
			for _, s := range ss {
				i, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				u = append(u, i)
			}
			data.updates = append(data.updates, u)
		}
	}
	return data
}

type Data struct {
	rules   map[int][]int
	updates [][]int
}

func Part1(data Data) int {
	sum := 0
	for _, u := range data.updates {
		correct := true
		before := make(map[int]bool)
		for _, i := range u {
			for _, r := range data.rules[i] {
				if before[r] {
					correct = false
					break
				}
			}
			if !correct {
				break
			}
			before[i] = true
		}
		if correct {
			sum += u[len(u)/2]
		}
	}
	return sum
}

func Part2(data Data) int {
	fmt.Println(data.rules)
	sum := 0
	for _, u := range data.updates {
		correct := true
		before := make(map[int]bool)
		for _, i := range u {
			for _, r := range data.rules[i] {
				if before[r] {
					correct = false
					break
				}
			}
			if !correct {
				break
			}
			before[i] = true
		}
		if !correct {
			// find the middle item if it were sorted
			for i := 0; i < len(u); i++ {
				count := 0
				for _, j := range data.rules[u[i]] {
					if slices.Contains(u, j) {
						count++
					}
				}
				if count == len(u)/2 {
					sum += u[i]
					break
				}
			}
		}
	}
	return sum
}
