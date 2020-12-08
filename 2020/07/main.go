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
	bags, err := ReadInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1 = ", Part1(bags, "shiny gold"))

	fmt.Println("Part 2 = ", Part2(bags, "shiny gold"))
}

func recursiveParent(bags map[string]Bag, bagName string, found map[string]bool) {
	for parentBagName := range bags[bagName].ParentBags {
		if _, ok := found[parentBagName]; !ok {
			recursiveParent(bags, parentBagName, found)
			found[parentBagName] = true
		}
	}
}

func Part1(bags map[string]Bag, bagName string) int {
	found := make(map[string]bool)
	recursiveParent(bags, bagName, found)
	return len(found)
}

func recursiveChild(bags map[string]Bag, bagName string) int {
	count := 0
	for subBagName, i := range bags[bagName].SubBags {
		count += i + i*recursiveChild(bags, subBagName)
	}
	return count
}

func Part2(bags map[string]Bag, bagName string) int {
	return recursiveChild(bags, bagName)
}

type Bag struct {
	SubBags    map[string]int
	ParentBags map[string]bool
}

// ReadInput reads the input file
func ReadInput(inputPath string) (map[string]Bag, error) {
	bags := make(map[string]Bag)

	// Compile regex
	r, err := regexp.Compile("(\\d+)\\s(\\w+\\s\\w+)")
	if err != nil {
		return nil, err
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rule := strings.Split(scanner.Text(), " bags contain ")
		bagName := rule[0]
		if _, ok := bags[bagName]; !ok {
			bags[bagName] = Bag{SubBags: make(map[string]int), ParentBags: make(map[string]bool)}
		}
		bag := bags[bagName]

		if rule[1] != "no other bags." {
			for _, match := range r.FindAllStringSubmatch(rule[1], -1) {
				subBagCount, err := strconv.Atoi(match[1])
				if err != nil {
					return nil, err
				}
				subBagName := match[2]
				bag.SubBags[subBagName] = subBagCount

				if _, ok := bags[subBagName]; !ok {
					bags[subBagName] = Bag{SubBags: make(map[string]int), ParentBags: make(map[string]bool)}
				}
				subBag := bags[subBagName]
				subBag.ParentBags[bagName] = true
			}
		}
	}
	return bags, nil
}
