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
		panic("could not open input")
	}
	s := bufio.NewScanner(f)
	encoding := processInput(s)
	noPriorTwoSum := partOne(encoding)
	fmt.Println("part one: ", noPriorTwoSum)
	fmt.Println("part two: ", partTwo(encoding, noPriorTwoSum))
}

func processInput(s *bufio.Scanner) []int {
	encoding := make([]int, 0, 1000)
	for s.Scan() {
		t := s.Text()
		val, err := strconv.ParseInt(t, 0, 64)
		if err != nil {
			panic("could not process input")
		}
		encoding = append(encoding, int(val))

	}
	return encoding
}

func partOne(e []int) int {
	preambleLength := 25
	sortedSet := make([]int, preambleLength)
	copy(sortedSet, e[:preambleLength])
	sort.Ints(sortedSet)

	nextToRemove := 0
	for _, sum := range e[preambleLength:] {
		if !twoSum(sortedSet, sum) {
			return sum
		}
		sortedSet = remove(sortedSet, e[nextToRemove])
		sortedSet = insertOrdered(sortedSet, sum)
		nextToRemove++
	}

	return 0
}

func twoSum(a []int, sum int) bool {
	left, right := 0, len(a)-1
	for a[left] < a[right] {
		if a[left]+a[right] < sum {
			left++
		} else if a[left]+a[right] > sum {
			right--
		} else {
			return true
		}
	}
	return false
}

func remove(s []int, v int) []int {
	index := sort.SearchInts(s, v)
	if index == len(s) || s[index] != v {
		panic("value not in array")
	}
	if index < len(s)-1 {
		copy(s[index:], s[index+1:])
	}
	s = s[:len(s)-1]
	return s
}

func insertOrdered(s []int, v int) []int {
	index := sort.SearchInts(s, v)
	s = append(s[:index], append([]int{v}, s[index:]...)...)
	return s
}

func partTwo(e []int, sum int) int {
	result := recurseContiguous(e, sum, 0, 1)
	return result
}

type indices struct {
	start int
	end   int
}

var seen = make(map[indices]bool)

func recurseContiguous(e []int, sum int, start int, end int) int {
	coordinates := indices{start: start, end: end}
	if start >= end || start > len(e) || end > len(e) || seen[coordinates] {
		return 0
	}
	seen[coordinates] = true
	if slice := e[start:end]; arraySum(slice) == sum && end-start > 1 {
		minMaxSum := Min(slice) + Max(slice)
		fmt.Println("min: ", Min(slice), " | max: ", Max(slice))
		return minMaxSum
	}
	if res := recurseContiguous(e, sum, start, end+1); res != 0 {
		return res
	}
	if res := recurseContiguous(e, sum, start+1, end); res != 0 {
		return res
	}
	return 0
}

func arraySum(a []int) (total int) {
	for _, val := range a {
		total += val
	}
	return
}

func Min(a []int) (minValue int) {
	return a[arrayMin(a, 0, len(a))]
}

func Max(a []int) (minValue int) {
	return a[arrayMax(a, 0, len(a))]
}

func arrayMin(a []int, start int, end int) (minIndex int) {
	minIndex = start
	for i := start + 1; i < end; i++ {
		if a[i] < a[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func arrayMax(a []int, start int, end int) (maxIndex int) {
	maxIndex = start
	for i := start + 1; i < end; i++ {
		if a[i] > a[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
