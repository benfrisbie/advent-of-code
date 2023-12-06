package main

import (
	"bufio"
	"fmt"
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

type Set struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID   int
	Sets []Set
}

// ReadInput reads the input file
func ReadInput(inputPath string) []Game {
	var data []Game
	// Open input file
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gameRegex := regexp.MustCompile(`Game (\d+)`)
	setRegex := regexp.MustCompile(`(\d+)\s(red|blue|green)`)

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		splitGame := strings.Split(s, ":")
		game := Game{}
		game.ID, err = strconv.Atoi(gameRegex.FindAllStringSubmatch(splitGame[0], -1)[0][1])
		if err != nil {
			panic(err)
		}
		for _, ss := range strings.Split(splitGame[1], ";") {
			set := Set{}
			for _, m := range setRegex.FindAllStringSubmatch(ss, -1) {
				switch m[2] {
				case "red":
					set.Red, _ = strconv.Atoi(m[1])
				case "green":
					set.Green, _ = strconv.Atoi(m[1])
				case "blue":
					set.Blue, _ = strconv.Atoi(m[1])
				}
			}
			game.Sets = append(game.Sets, set)
		}
		data = append(data, game)
	}
	return data
}

func Part1(data []Game) int {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	sum := 0

	gameValid := true
	for _, game := range data {
		gameValid = true
		for _, set := range game.Sets {
			if set.Red > maxRed || set.Green > maxGreen || set.Blue > maxBlue {
				gameValid = false
				break
			}
		}
		if gameValid {
			sum += game.ID
		}
	}

	return sum
}

func Part2(data []Game) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0
	sum := 0

	for _, game := range data {
		maxRed = 0
		maxGreen = 0
		maxBlue = 0
		for _, set := range game.Sets {
			if set.Red > maxRed {
				maxRed = set.Red
			}
			if set.Green > maxGreen {
				maxGreen = set.Green
			}
			if set.Blue > maxBlue {
				maxBlue = set.Blue
			}
		}
		sum += maxRed * maxGreen * maxBlue
	}

	return sum
}
