package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

type Card struct {
	ID             int
	WinningNumbers map[int]bool
	PresentNumbers []int
}

func (card Card) CountOfWinningNumbers() int {
	count := 0
	for _, num := range card.PresentNumbers {
		if card.WinningNumbers[num] {
			count++
		}
	}
	return count
}

func ReadInput(inputPath string) []Card {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data []Card

	re := regexp.MustCompile(`(\d+)`)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(strings.Split(s, ":")[1], "|")
		card := Card{ID: i, WinningNumbers: map[int]bool{}}
		for _, m := range re.FindAllStringSubmatch(split[0], -1) {
			num, err := strconv.Atoi(m[1])
			if err != nil {
				panic(err)
			}
			card.WinningNumbers[num] = true
		}
		for _, m := range re.FindAllStringSubmatch(split[1], -1) {
			num, err := strconv.Atoi(m[1])
			if err != nil {
				panic(err)
			}
			card.PresentNumbers = append(card.PresentNumbers, num)
		}
		data = append(data, card)
		i++
	}

	return data
}

func Part1(data []Card) int {
	sum := 0

	for _, card := range data {
		sum += int(math.Pow(2, float64(card.CountOfWinningNumbers()-1)))
	}

	return sum
}

func Part2(data []Card) int {
	cardCount := map[int]int{}

	for i := range data {
		cardCount[i]++
		winningCount := data[i].CountOfWinningNumbers()
		for j := i + 1; j < i+winningCount+1; j++ {
			cardCount[j] += cardCount[i]
		}
	}

	sum := 0
	for _, v := range cardCount {
		sum += v
	}
	return sum
}
