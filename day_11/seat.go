package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	l := ProcessInput()
	occupied := partOne(l)
	fmt.Println("part one")
	fmt.Printf("After stable state, %d seats occupied.\n", occupied)
	isModified = true
	fmt.Println("part two")
	occupied = partTwo(l)
	fmt.Printf("After stable state, %d seats occupied.\n", occupied)
}

type layout struct {
	seats [][]rune
}

var isModified bool

func partOne(l *layout) int {
	for {
		next := l.NextState()
		if l.IsEqual(next) {
			return l.Occupied()
		}
		l = next
	}
}

func partTwo(l *layout) int {
	for {
		next := l.NextState()
		if l.IsEqual(next) {
			return l.Occupied()
		}
		l = next
	}
}

func (l *layout) PrettyPrint() {
	for _, row := range l.seats {
		fmt.Println(string(row))
	}
}

func (l *layout) NextState() *layout {
	retLayout := &layout{seats: [][]rune{}}
	for indrow, row := range l.seats {
		retRow := make([]rune, len(row))
		for indspace := range row {
			retRow[indspace] = nextStateSeat(l.seats, indrow, indspace)
		}
		retLayout.seats = append(retLayout.seats, retRow)
	}
	return retLayout
}

func nextStateSeat(seats [][]rune, row int, col int) rune {
	var tooManySeats int
	if isModified {
		tooManySeats = 5
	} else {
		tooManySeats = 4
	}
	position := seats[row][col]
	if position == rune(0x2e) {
		return rune(0x2e)
	}

	var adjacentOccupied int
	if isModified {
		adjacentOccupied = numAdjacentOccupiedModified(seats, row, col)
	} else {
		adjacentOccupied = numAdjacentOccupied(seats, row, col)
	}

	if position == rune(0x4c) {
		if adjacentOccupied > 0 {
			return rune(0x4c)
		} else {
			return rune(0x23)
		}
	} else {
		if adjacentOccupied >= tooManySeats {
			return rune(0x4c)
		} else {
			return rune(0x23)
		}

	}
}

func numAdjacentOccupiedModified(seats [][]rune, row int, col int) int {
	adjacentOccupied := 0
	var steps [][]int = [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	for _, step := range steps {
		for i := 1; i*step[1]+row >= 0 && i*step[1]+row < len(seats) && i*step[0]+col >= 0 && i*step[0]+col < len(seats[0]); i++ {
			space := seats[i*step[1]+row][i*step[0]+col]
			if space == rune(0x23) {
				adjacentOccupied++
				break
			} else if space == rune(0x4c) {
				break
			}
		}
	}
	return adjacentOccupied
}

func numAdjacentOccupied(seats [][]rune, row int, col int) int {
	adjacentOccupied := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if (row+i >= 0 && row+i < len(seats) && col+j >= 0 && col+j < len(seats[0])) && seats[row+i][col+j] == rune(0x23) {
				adjacentOccupied++
			}
		}
	}
	return adjacentOccupied
}

func (l1 *layout) IsEqual(l2 *layout) bool {
	for indrow, row1 := range l1.seats {
		for indspace, space1 := range row1 {
			if space1 != l2.seats[indrow][indspace] {
				return false
			}
		}
	}
	return true
}

func (l *layout) Occupied() int {
	occupied := 0
	for _, row := range l.seats {
		for _, space := range row {
			if space == rune(35) {
				occupied += 1
			}

		}
	}
	return occupied

}

func ProcessInput() *layout {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("input")
	}
	s := bufio.NewScanner(f)

	l := &layout{seats: [][]rune{}}
	for s.Scan() {
		row := []rune{}
		for _, val := range s.Text() {
			if val != rune(0x2e) && val != rune(0x23) && val != rune(0x4c) {
				fmt.Printf("unexpected value: %s [%x] \n", string(val), val)
				panic("exit on input bad")
			}
			row = append(row, val)
		}
		l.seats = append(l.seats, row)
	}
	return l
}
