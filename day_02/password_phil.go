package main

import (
    "os"
    "bufio"
    "strconv"
    "fmt"
    "strings"
)


func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic("file not read")
    }
    s := bufio.NewScanner(f)
    valid := 0
    for s.Scan() {
        t := s.Text()
        if isValidIndex(t) {
            valid ++
        }
    }
    fmt.Printf("num valid: %d\n", valid)
}

func isValidCount(text string) bool {
    sections := strings.Split(text, " ")
    letter := sections[1][0]
    bounds := strings.Split(sections[0], "-")
    password := sections[2]
    lowerbound, err := strconv.ParseInt(bounds[0], 0, 64)
    if err != nil {
        panic("whoops")
    }
    upperbound, err := strconv.ParseInt(bounds[1], 0, 64)
    if err != nil {
        panic("whoops")
    }
    count := 0
    for _, r := range(password) {
        if rune(letter) == r {
            count ++
        }
    }
    isvalid := count >= int(lowerbound) && count <= int(upperbound)
    // fmt.Printf("orig string [ %s ], is valid = %t\n", text, isvalid)
    return isvalid
}

func isValidIndex(text string) bool {
    sections := strings.Split(text, " ")
    letter := sections[1][0]
    bounds := strings.Split(sections[0], "-")
    password := sections[2]
    lowerindex, err := strconv.ParseInt(bounds[0], 0, 64)
    if err != nil {
        panic("whoops")
    }
    upperindex, err := strconv.ParseInt(bounds[1], 0, 64)
    if err != nil {
        panic("whoops")
    }
    if int(lowerindex) - 1 >= len(password) {
        return false
    }
    presentAtLow := letter == password[lowerindex - 1]
    presentAtHigh := false
    if int(upperindex) - 1 < len(password) {
        presentAtHigh = letter == password[upperindex - 1]
    } 
    isvalid := presentAtLow && !presentAtHigh || !presentAtLow && presentAtHigh
    // fmt.Printf("orig string [ %s ], is valid = %t\n", text, isvalid)
    return isvalid
}

