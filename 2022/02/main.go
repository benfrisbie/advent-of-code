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
	totalScore := 0
	for _, p := range data {
		totalScore += ScoreRound1(p)
	}
	fmt.Printf("Part 1 = %d\n", totalScore)

	// Part 2
	totalScore = 0
	for _, p := range data {
		totalScore += ScoreRound2(p)
	}
	fmt.Printf("Part 2 = %d\n", totalScore)
}

type Input int

const (
	Rock Input = iota
	Paper
	Scissor
)

const (
	Lose Input = iota
	Tie
	Win
)

type Pair struct {
	First  Input
	Second Input
}

func ScoreRound1(p Pair) int {
	var score int
	switch p.Second {
	case Rock:
		score = 1
		switch p.First {
		case Rock:
			score += 3
		case Paper:
			score += 0
		case Scissor:
			score += 6
		}
	case Paper:
		score = 2
		switch p.First {
		case Rock:
			score += 6
		case Paper:
			score += 3
		case Scissor:
			score += 0
		}
	case Scissor:
		score = 3
		switch p.First {
		case Rock:
			score += 0
		case Paper:
			score += 6
		case Scissor:
			score += 3
		}
	}
	return score
}

func ScoreRound2(p Pair) int {
	var score int
	switch p.Second {
	case Lose:
		score = 0
		switch p.First {
		case Rock:
			score += 3
		case Paper:
			score += 1
		case Scissor:
			score += 2
		}
	case Tie:
		score = 3
		switch p.First {
		case Rock:
			score += 1
		case Paper:
			score += 2
		case Scissor:
			score += 3
		}
	case Win:
		score = 6
		switch p.First {
		case Rock:
			score += 2
		case Paper:
			score += 3
		case Scissor:
			score += 1
		}
	}
	return score
}

// ReadInput reads the input file
func ReadInput(inputPath string) ([]Pair, error) {
	// Open input file
	var data []Pair
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var p Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		// extract opponents hand symbol
		switch split[0] {
		case "A":
			p.First = Rock
		case "B":
			p.First = Paper
		case "C":
			p.First = Scissor
		default:
			panic("unknown hand symbol")
		}
		// extract my hand symbol or if we should win/lose/tie
		switch split[1] {
		case "X":
			p.Second = Rock
		case "Y":
			p.Second = Paper
		case "Z":
			p.Second = Scissor
		default:
			panic("unknown hand symbol")
		}
		data = append(data, p)
	}
	return data, nil
}
