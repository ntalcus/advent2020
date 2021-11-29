package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("uh oh")
	}
	s := bufio.NewScanner(f)

	fmt.Println("part one")
	partOne(s)
	f.Seek(0, 0)
	s = bufio.NewScanner(f)
	fmt.Println("part two")
	partTwo(s)
}

type bag struct {
	color    string
	quantity int
}

// containedBy is a map
// key := bag color
// value := array of bags that contain this color (color + the quantities of the current color held)
var containedBy = make(map[string][]bag)

func partOne(s *bufio.Scanner) {
	targetColor := "shiny gold"
	for s.Scan() {
		t := s.Text()
		bagColor, contains, quantities := processLine(t)
		for ind, containedBag := range contains {
			if containedBy[containedBag] == nil {
				containedBy[containedBag] = []bag{{color: bagColor, quantity: quantities[ind]}}
			} else {
				containedBy[containedBag] = append(containedBy[containedBag], bag{color: bagColor, quantity: quantities[ind]})
			}
		}
	}
	traverseGoldHolders(containedBy[targetColor])
	fmt.Printf("%d bags can hold at least one gold.\n", len(seen))
	fmt.Printf("%d total bags. \n", len(containedBy))
}

var seen = make(map[string]bool)

func traverseGoldHolders(bags []bag) {
	for _, bag := range bags {
		if !seen[bag.color] {
			seen[bag.color] = true
			traverseGoldHolders(containedBy[bag.color])
		}
	}
}

func processLine(line string) (bagColor string, contains []string, quantities []int) {
	components := strings.Split(line, " contain ")
	bagColor = strings.TrimSpace(strings.Replace(components[0], "bags", "", -1))
	if components[1] == "no other bags." {
		return
	}
	containedBags := strings.Split(components[1][:len(components[1])-1], ", ")
	for _, bagStr := range containedBags {
		details := strings.SplitAfterN(bagStr, " ", 2)
		number, err := strconv.ParseInt(strings.TrimSpace(details[0]), 0, 64)
		if err != nil {
			panic("non number in contain value")
		}
		containsCleaned := strings.Replace(details[1], "bags", "", -1)
		containsCleaned = strings.TrimSpace(strings.Replace(containsCleaned, "bag", "", -1))
		contains = append(contains, containsCleaned)
		quantities = append(quantities, int(number))
	}
	return
}

// whichContains is a map
// key := bag color
// value := array of bags that are contained by this color (color + the quantities of the current color held)
var whichContains = make(map[string][]bag)
var seenPartTwo = make(map[string]int)

func partTwo(s *bufio.Scanner) {
	targetColor := "shiny gold"
	for s.Scan() {
		t := s.Text()
		bagColor, contains, quantities := processLine(t)
		containsBags := make([]bag, len(contains))
		for ind, containedBag := range contains {
			containsBags[ind] = bag{color: containedBag, quantity: quantities[ind]}
		}
		whichContains[bagColor] = containsBags
	}
	numHeld := traverseHeldByGold(whichContains[targetColor])
	fmt.Printf("a gold bag holds %d bags\n", numHeld)
}

func traverseHeldByGold(bags []bag) (total int) {
	for _, bag := range bags {
		if seenPartTwo[bag.color] == 0 {
			seenPartTwo[bag.color] = traverseHeldByGold(whichContains[bag.color])
		}
		total += (1 + seenPartTwo[bag.color]) * bag.quantity
	}
	return
}
