package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

type History []int

func (h History) Predict() int {
	done := true
	for _, v := range h {
		if v != 0 {
			done = false
			break
		}
	}
	if done {
		return 0
	}

	diffHistory := make(History, len(h)-1)
	for i := 0; i < len(diffHistory); i++ {
		diffHistory[i] = h[i+1] - h[i]
	}
	return diffHistory.Predict() + h[len(h)-1]
}

func ReadInput(inputPath string) []History {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data []History

	re := regexp.MustCompile(`(-?\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		var h History
		for _, match := range re.FindAllStringSubmatch(s, -1) {
			i, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			h = append(h, i)
		}
		data = append(data, h)
	}

	return data
}

func Part1(data []History) int {
	sum := 0
	for _, h := range data {
		sum += h.Predict()
	}
	return sum
}

func Part2(data []History) int {
	sum := 0
	for _, h := range data {
		slices.Reverse(h)
		sum += h.Predict()
	}
	return sum
}
