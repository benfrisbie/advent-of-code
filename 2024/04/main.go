package main

import (
	"bufio"
	"fmt"
	"os"
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
		data.grid = append(data.grid, []rune(s))
	}
	return data
}

type Data struct {
	grid [][]rune
}

func searchForWord(data Data, i, j int, word string) int {
	if data.grid[i][j] != rune(word[0]) {
		return 0
	}
	if len(word) == 1 {
		return 1
	}
	count := 0

	var a, b int
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			for l := 1; l < len(word); l++ {
				a = i + x*l
				b = j + y*l
				if a < 0 || a >= len(data.grid) {
					break
				}
				if b < 0 || b >= len(data.grid[i]) {
					break
				}
				if data.grid[a][b] != rune(word[l]) {
					break
				}
				if l == len(word)-1 {
					count++
				}
			}
		}
	}
	return count
}

func Part1(data Data) int {
	sum := 0
	for i := range data.grid {
		for j := range data.grid[i] {
			sum += searchForWord(data, i, j, "XMAS")
		}
	}
	return sum
}

func Part2(data Data) int {
	sum := 0
	for i := range data.grid {
		for j := range data.grid[i] {
			if data.grid[i][j] == 'A' {
				if i-1 < 0 || i+1 >= len(data.grid) || j-1 < 0 || j+1 >= len(data.grid[i]) {
					continue
				}
				ul := data.grid[i-1][j-1]
				ur := data.grid[i-1][j+1]
				dl := data.grid[i+1][j-1]
				dr := data.grid[i+1][j+1]
				if ((ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M')) && ((ur == 'M' && dl == 'S') || (ur == 'S' && dl == 'M')) {
					sum++
				}
			}
		}
	}
	return sum
}
