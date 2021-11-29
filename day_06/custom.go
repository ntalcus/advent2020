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
	fmt.Println("Part One")
	partOne(s)
	offset, err := f.Seek(0, 0)
	if offset != 0 || err != nil {
		panic("uh oh")
	}
	s = bufio.NewScanner(f)
	fmt.Println("Part Two")
	partTwo(s)
}

func partOne(s *bufio.Scanner) {
	count := 0
	group := map[rune]bool{}
	for s.Scan() {
		t := s.Text()
		if len(t) == 0 {
			for range group {
				count++
			}
			group = map[rune]bool{}
		}
		for _, char := range t {
			group[char] = true
		}
	}
	for range group {
		count++
	}
	fmt.Println(count)
}

func partTwo(s *bufio.Scanner) {
	count := 0
	group := map[rune]int{}
	memberCount := 0
	for s.Scan() {
		t := s.Text()
		if len(t) > 0 {
			for _, char := range t {
				group[char] += 1
			}
			memberCount++
		}
		if len(t) == 0 {
			for _, val := range group {
				if val == memberCount {
					count++
				}
			}
			group = map[rune]int{}
			memberCount = 0
		}
	}
	for _, val := range group {
		if val == memberCount {
			count++
		}
	}
	fmt.Println(count)

}
