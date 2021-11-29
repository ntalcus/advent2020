package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("could not open input")
	}
	s := bufio.NewScanner(f)
	instr := parseInstructions(s)
	fmt.Printf("part one accumulator: %d\n", runInstructions(instr))
	var acc int
	for _, faultyInd := range maybeFaulty {
		flipInstr(instr, faultyInd)
		acc = runInstructions(instr)
		if ranToCompletion {
			fmt.Println("ran to completion")
			break
		}
		flipInstr(instr, faultyInd)
		resetSeen(instr)
	}
	fmt.Printf("part two accumulator: %d\n", acc)
}

func flipInstr(instr []*instruction, faultyInd int) {
	if instr[faultyInd].operation == nop {
		instr[faultyInd].operation = jmp
	} else if instr[faultyInd].operation == jmp {
		instr[faultyInd].operation = nop
	}
}

func resetSeen(instr []*instruction) {
	for _, i := range instr {
		i.seen = false
	}
}

type instruction struct {
	operation int
	value     int
	seen      bool
}

const (
	nop = iota
	jmp
	acc
)

var maybeFaulty = make([]int, 0, 1000)
var ranToCompletion bool

func parseInstructions(s *bufio.Scanner) []*instruction {
	instr := make([]*instruction, 0, 1000)
	for s.Scan() {
		i := processInstruction(s.Text())
		if i.operation == nop || i.operation == jmp {
			maybeFaulty = append(maybeFaulty, len(instr))
		}
		instr = append(instr, i)
	}
	return instr
}

func processInstruction(line string) (i *instruction) {
	switch op := line[:3]; op {
	case "nop":
		i = &instruction{operation: nop}
	case "jmp":
		i = &instruction{operation: jmp}
	case "acc":
		i = &instruction{operation: acc}
	}
	val, err := strconv.ParseInt(line[4:], 0, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot interpret instruction: %s\n", line))
	}
	i.value = int(val)
	return i
}

func runInstructions(instr []*instruction) int {
	instructionCount := 0
	accumulator := 0
	for {
		if instructionCount >= len(instr) {
			ranToCompletion = true
			return accumulator
		}
		ci := instr[instructionCount]
		if ci.seen {
			return accumulator
		}
		ci.seen = true
		switch ci.operation {
		case jmp:
			instructionCount += ci.value
		case acc:
			accumulator += ci.value
			instructionCount++
		case nop:
			instructionCount++
		}
	}
}
