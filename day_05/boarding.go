package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("file not read")
	}
	s := bufio.NewScanner(f)
	partOne(s)
	offset, err := f.Seek(0, 0)
	if offset != 0 || err != nil {
		panic("uh oh")
	}
	s = bufio.NewScanner(f)
	partTwo(s)
}

func partOne(s *bufio.Scanner) {
	fmt.Println("Part One")
	largestID := -1
	for s.Scan() {
		t := s.Text()
		id := 0
		for ind, char := range t {
			if char == rune(66) || char == rune(82) {
				id += power(2, 9-ind)
			}

		}
		if id > largestID {
			largestID = id
		}
	}
	fmt.Println(largestID)
}

func partTwo(s *bufio.Scanner) {
	greatestID := 843
	fmt.Println("Part Two")
	tickets := make([]int, greatestID)
	lowestID := greatestID + 1
	for s.Scan() {
		t := s.Text()
		id := 0
		for ind, char := range t {
			if char == rune(66) || char == rune(82) {
				id += power(2, 9-ind)
			}

		}
		tickets[id] = id
		if id < lowestID {
			lowestID = id
		}
	}
	for i := lowestID; i < greatestID; i++ {
		if i != tickets[i] {
			fmt.Println(i, tickets[i])
		}
	}
}

func power(base int, exponent int) int {
	// do not use for negatives
	result := 1
	for exponent > 0 {
		result *= base
		exponent--
	}
	return result
}
