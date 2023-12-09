package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

type Node struct {
	Left  string
	Right string
}

type Map struct {
	Instructions string
	// StartingNode string
	Nodes map[string]Node
}

func ReadInput(inputPath string) Map {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := Map{Nodes: make(map[string]Node)}

	re := regexp.MustCompile(`(\w{3})\s+=\s+\((\w{3}),\s+(\w{3})\)`)

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if i == 0 {
			data.Instructions = s
			scanner.Scan()
			i++
		} else {
			for _, match := range re.FindAllStringSubmatch(s, -1) {
				data.Nodes[match[1]] = Node{Left: match[2], Right: match[3]}
			}
		}
	}

	return data
}

func Part1(data Map) int {
	cur := "AAA"
	i := 0
	for ; cur != "ZZZ"; i++ {
		// fmt.Println(cur)
		switch data.Instructions[i%len(data.Instructions)] {
		case 'L':
			cur = data.Nodes[cur].Left
		case 'R':
			cur = data.Nodes[cur].Right
		}
	}
	return i
}

func Part2(data Map) int {
	var cur string
	var cycleLengths []int
	for node := range data.Nodes {
		if node[2] == 'A' {
			cur = node
			for i := 0; ; i++ {
				if cur[2] == 'Z' {
					cycleLengths = append(cycleLengths, i)
					break
				}
				switch data.Instructions[i%len(data.Instructions)] {
				case 'L':
					cur = data.Nodes[cur].Left
				case 'R':
					cur = data.Nodes[cur].Right
				}
			}
		}
	}

	return LeastCommonMultiple(cycleLengths)
}

// GreatestCommonDivisor calculates the greatest common divisor of two integers
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func leastCommonMultiple(a, b int) int {
	return a * b / GreatestCommonDivisor(a, b)
}

// LeastCommonMultiple calculates the least common multiple of a slice of integers
func LeastCommonMultiple(ints []int) int {
	lcm := ints[0]
	for _, num := range ints[1:] {
		lcm = leastCommonMultiple(lcm, num)
	}
	return lcm
}
