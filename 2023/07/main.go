package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	data := ReadInput("input.txt")

	fmt.Println("Part1:", Part1(data))

	fmt.Println("Part2:", Part2(data))
}

var cardMappingPart1 = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardMappingPart2 = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Hand struct {
	Cards [5]rune
	Bid   int
}

func (h Hand) String() string {
	return fmt.Sprintf("{[%v%v%v%v%v] %v}", string(h.Cards[0]), string(h.Cards[1]), string(h.Cards[2]), string(h.Cards[3]), string(h.Cards[4]), h.Bid)
}

func (h Hand) Type(jokersWild bool) int {
	jokers := 0
	fiveOfAKind := false
	fourOfAKind := false
	threeOfAKind := false
	pairs := 0
	cardMap := map[rune]int{}
	for _, card := range h.Cards {
		if jokersWild && card == 'J' {
			jokers++
			continue
		}
		cardMap[card]++
		if cardMap[card] == 5 {
			fiveOfAKind = true
		} else if cardMap[card] == 4 {
			fourOfAKind = true
		} else if cardMap[card] == 3 {
			threeOfAKind = true
			pairs--
		} else if cardMap[card] == 2 {
			pairs++
		}
	}
	if fiveOfAKind {
		return 6
	} else if fourOfAKind {
		if jokers > 0 {
			return 6
		}
		return 5
	} else if threeOfAKind && pairs == 1 {
		return 4
	} else if threeOfAKind {
		if jokers > 0 {
			return 4 + jokers
		}
		return 3
	} else if pairs == 2 {
		if jokers > 0 {
			return 4
		}
		return 2
	} else if pairs == 1 {
		if jokers > 0 {
			if jokers == 1 {
				return 3
			}
			if jokers == 2 {
				return 5
			}
			if jokers == 3 {
				return 6
			}
		}
		return 1
	} else {
		if jokers > 0 {
			if jokers == 1 {
				return 1
			}
			if jokers == 2 {
				return 3
			}
			if jokers == 3 {
				return 5
			}
			if jokers >= 4 {
				return 6
			}
		}
		return 0
	}
}

type Hands []Hand

func (hands Hands) Len() int {
	return len(hands)
}

func (hands Hands) Less(i, j int) bool {
	ti := hands[i].Type(false)
	tj := hands[j].Type(false)
	if ti != tj {
		return ti < tj
	}
	for k := 0; k < 5; k++ {
		ti := cardMappingPart1[hands[i].Cards[k]]
		tj := cardMappingPart1[hands[j].Cards[k]]
		if ti != tj {
			return ti < tj
		}
	}
	return true
}

func (hands Hands) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

type HandsPart2 []Hand

func (hands HandsPart2) Len() int {
	return len(hands)
}

func (hands HandsPart2) Less(i, j int) bool {
	ti := hands[i].Type(true)
	tj := hands[j].Type(true)
	if ti != tj {
		return ti < tj
	}
	for k := 0; k < 5; k++ {
		ti := cardMappingPart2[hands[i].Cards[k]]
		tj := cardMappingPart2[hands[j].Cards[k]]
		if ti != tj {
			return ti < tj
		}
	}
	return true
}

func (hands HandsPart2) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

func ReadInput(inputPath string) Hands {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data Hands

	re := regexp.MustCompile(`(\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		hand := Hand{}
		for i := 0; i < 5; i++ {
			hand.Cards[i] = rune(s[i])
		}
		hand.Bid, err = strconv.Atoi(re.FindString(s[5:]))
		if err != nil {
			panic(err)
		}
		data = append(data, hand)
	}

	return data
}

func Part1(data Hands) int {
	sum := 0
	sort.Sort(data)
	for i, hand := range data {
		sum += hand.Bid * (i + 1)
	}
	return sum
}

func Part2(data Hands) int {
	sum := 0
	handsPart2 := HandsPart2(data)
	sort.Sort(handsPart2)
	for i, hand := range handsPart2 {
		sum += hand.Bid * (i + 1)
	}
	return sum
}
