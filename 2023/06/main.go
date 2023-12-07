package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

type Race struct {
	Time           int
	RecordDistance int
}

func (r Race) CountIntegerSolutions() int {
	a := float64(1)
	b := float64(-r.Time)
	c := float64(r.RecordDistance)

	lower := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2
	lowerint := int(math.Ceil(lower))
	if lower == float64(int(lower)) {
		lowerint++
	}

	upper := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2
	upperint := int(math.Floor(upper))
	if upper == float64(int(upper)) {
		upperint--
	}

	return upperint - lowerint + 1
}

func ReadInput(inputPath string) []Race {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data []Race

	i := 0
	re := regexp.MustCompile(`(\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		for j, match := range re.FindAllStringSubmatch(s, -1) {
			if i == 0 {
				time, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				data = append(data, Race{Time: time})
			} else {
				recordDistance, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				data[j].RecordDistance = recordDistance
			}
		}
		i++
	}

	return data
}

func Part1(data []Race) int {
	answer := 1
	for _, race := range data {
		answer *= race.CountIntegerSolutions()
	}
	return answer
}

func Part2(data []Race) int {
	combinedTimeString := ""
	combinedDistanceString := ""

	for _, race := range data {
		combinedTimeString += fmt.Sprintf("%d", race.Time)
		combinedDistanceString += fmt.Sprintf("%d", race.RecordDistance)
	}
	combinedTime, err := strconv.Atoi(combinedTimeString)
	if err != nil {
		panic(err)
	}
	combinedDistance, err := strconv.Atoi(combinedDistanceString)
	if err != nil {
		panic(err)
	}

	race := Race{Time: combinedTime, RecordDistance: combinedDistance}
	return race.CountIntegerSolutions()
}
