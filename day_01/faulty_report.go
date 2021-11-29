package main

import (
    "os"
    "bufio"
    "strconv"
    "fmt"
)


func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic("file not read")
    }
    s := bufio.NewScanner(f)
    var sum int64 = 2020
    seen := make(map[int64]bool)
    for s.Scan() {
        t := s.Text()
        num, err := strconv.ParseInt(t, 0, 64)
        if err != nil {
            panic("non convertable int")
        }
        if seen[sum - num] {
            fmt.Println((sum - num) * num)
            return
        }
        seen[num] = true
    }
}
