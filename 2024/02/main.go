package main

import (
	"bufio"
	"fmt"
	"os"
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
	var data Data
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
		ss := strings.Split(s, " ")
		report := make([]int, len(ss))
		for i, rs := range ss {
			n, err := strconv.Atoi(rs)
			if err != nil {
				panic(err)
			}
			report[i] = n
		}
		data.reports = append(data.reports, report)
	}
	return data
}

type Data struct {
	reports [][]int
}

func Part1(data Data) int {
	sum := 0
	for _, report := range data.reports {
		increasing := report[1] > report[0]
		safe := true
		for i := 1; i < len(report); i++ {
			if increasing && report[i] < report[i-1] {
				safe = false
				break
			}
			if !increasing && report[i] > report[i-1] {
				safe = false
				break
			}
			diff := report[i] - report[i-1]
			if diff == 0 || diff > 3 || diff < -3 {
				safe = false
				break
			}
		}
		if safe {
			sum++
		}
	}
	return sum
}

func Part2(data Data) int {
	sum := 0
	for _, r := range data.reports {
		fmt.Println("r =", r)
		for j := -1; j < len(r); j++ {
			report := make([]int, len(r))
			copy(report, r)
			if j >= 0 {
				report = append(report[:j], report[j+1:]...)
			}
			fmt.Println("report =", report)
			increasing := report[1] > report[0]
			safe := true
			for i := 1; i < len(report); i++ {
				if increasing && report[i] < report[i-1] {
					safe = false
					break
				}
				if !increasing && report[i] > report[i-1] {
					safe = false
					break
				}
				diff := report[i] - report[i-1]
				if diff == 0 || diff > 3 || diff < -3 {
					safe = false
					break
				}
			}
			if safe {
				sum++
				break
			}
		}
	}
	return sum
}
