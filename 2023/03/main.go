package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

type Node struct {
	Number   *int
	Symbol   *rune
	Adjacent []Node
}

func (n Node) String() string {
	if n.Number != nil {
		return fmt.Sprintf("%d", *n.Number)
	} else if n.Symbol != nil {
		return fmt.Sprintf("%c", *n.Symbol)
	}
	return "."
}

// ReadInput reads the input file
func ReadInput(inputPath string) []Node {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]int
	var nodes []Node
	nodeIndex := 0

	numberInProgress := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		row := []int{}
		for i := range s {
			r := rune(s[i])
			if unicode.IsDigit(r) {
				numberInProgress += string(r)
				row = append(row, nodeIndex)
			} else {
				if numberInProgress != "" {
					nodeIndex++
					num, err := strconv.Atoi(numberInProgress)
					if err != nil {
						panic(err)
					}
					nodes = append(nodes, Node{Number: &num})
					numberInProgress = ""
				}

				if r == '.' {
					nodes = append(nodes, Node{})
				} else {
					nodes = append(nodes, Node{Symbol: &r})
				}

				row = append(row, nodeIndex)
				nodeIndex++
			}
		}
		if numberInProgress != "" {
			nodeIndex++
			num, err := strconv.Atoi(numberInProgress)
			if err != nil {
				panic(err)
			}
			nodes = append(nodes, Node{Number: &num})
			numberInProgress = ""
		}
		grid = append(grid, row)
	}

	rows := len(grid)
	columns := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			adjacentIndices := map[int]bool{}
			if i > 0 && j > 0 {
				adjacentIndices[grid[i-1][j-1]] = true
			}
			if i > 0 {
				adjacentIndices[grid[i-1][j]] = true
			}
			if i > 0 && j < columns-1 {
				adjacentIndices[grid[i-1][j+1]] = true
			}
			if j > 0 {
				adjacentIndices[grid[i][j-1]] = true
			}
			if j < columns-1 {
				adjacentIndices[grid[i][j+1]] = true
			}
			if i < rows-1 && j > 0 {
				adjacentIndices[grid[i+1][j-1]] = true
			}
			if i < rows-1 {
				adjacentIndices[grid[i+1][j]] = true
			}
			if i < rows-1 && j < columns-1 {
				adjacentIndices[grid[i+1][j+1]] = true
			}

			for index := range adjacentIndices {
				nodes[grid[i][j]].Adjacent = append(nodes[grid[i][j]].Adjacent, nodes[index])
			}
		}
	}

	return nodes
}

func Part1(data []Node) int {
	sum := 0

	for node := range data {
		if data[node].Number != nil {
			for _, adjacent := range data[node].Adjacent {
				if adjacent.Symbol != nil {
					sum += *data[node].Number
					break
				}
			}
		}
	}

	return sum
}

func Part2(data []Node) int {
	sum := 0
	adjacentNumbers := []int{}

	for _, node := range data {
		if node.Symbol != nil && *node.Symbol == '*' {
			for _, adjacent := range node.Adjacent {
				if adjacent.Number != nil {
					adjacentNumbers = append(adjacentNumbers, *adjacent.Number)
				}
			}
			if len(adjacentNumbers) == 2 {
				sum += adjacentNumbers[0] * adjacentNumbers[1]
			}
			adjacentNumbers = nil
		}
	}

	return sum
}
