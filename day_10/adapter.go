package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic("file")
	}
	s := bufio.NewScanner(f)
	a := processInput(s)
	partOne(a)
	a = append([]int{0}, a...)
	partTwo(a)
}

func processInput(s *bufio.Scanner) (a []int) {
	for s.Scan() {
		line := s.Text()
		lineVal, err := strconv.ParseInt(line, 0, 64)
		if err != nil {
			panic("input error")
		}
		a = append(a, int(lineVal))
	}
	sort.Ints(a)
	a = append(a, a[len(a)-1]+3)
	return a
}

var seenDiffs = []int{0, 0, 0, 0}

func partOne(a []int) {
	last := 0
	for _, val := range a {
		diff := val - last
		last = val
		seenDiffs[diff]++
	}
	fmt.Println(seenDiffs[1] * seenDiffs[3])
}

func partTwo(a []int) {
	visited[94] = 1
	fmt.Println(recurse(a, 0))
}

var visited = make([]int64, 95)

func recurse(a []int, node int) int64 {
	if node == len(a)-1 {
		return visited[len(a)-1]
	} else if node > len(a) {
		return 0
	}
	if visited[node] != 0 {
		return visited[node]
	}
	var total int64 = 0

	if node+1 < len(a) && (a[node+1] == a[node]+1 || a[node+1] == a[node]+2 || a[node+1] == a[node]+3) {
		total += recurse(a, node+1)
	}
	if node+2 < len(a) && (a[node+2] == a[node]+2 || a[node+2] == a[node]+3) {
		total += recurse(a, node+2)
	}
	if node+3 < len(a) && a[node+3] == a[node]+3 {
		total += recurse(a, node+3)
	}
	visited[node] = total
	return total
}
