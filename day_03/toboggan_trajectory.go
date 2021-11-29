package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    mainPartOne()
    mainPartTwo()
}


func mainPartOne() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic("file not read")
    }
    s := bufio.NewScanner(f)
    hits := 0
    positionX := 0
    // positionY := 0 
    xDelta := 3
    for s.Scan() {
        t := s.Text()
        if string(t[positionX]) == ("#") {
            hits ++
        }
        positionX = (positionX + xDelta) % len(t)
    }
    fmt.Printf("num hit: %d\n", hits)
}


func mainPartTwo() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic("file not read")
    }
    s := bufio.NewScanner(f)
    hits := []int{0, 0, 0, 0, 0}
    positionX := []int{0, 0, 0, 0, 0}
    xDelta := []int{1, 3, 5 ,7, 1}
    yDelta := []int{1, 1, 1, 1, 2}
    line := 0
    for s.Scan() {
        t := s.Text()
        for ind := range(hits) {
            if line % yDelta[ind] != 0 {
                continue
            }
            if string(t[positionX[ind]]) == "#" {
                hits[ind] += 1
            }
            positionX[ind] = (positionX[ind] + xDelta[ind]) % len(t)
        }
        line ++
    }
    fmt.Println(hits)
    fmt.Printf("num hit multiplier: %d\n", hits[0] * hits[1] * hits[2] * hits[3] * hits[4])
}

