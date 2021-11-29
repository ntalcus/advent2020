package main

import (
    "os"
    "bufio"
    "strconv"
    "fmt"
    "sort"
)


func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic("file not read")
    }
    s := bufio.NewScanner(f)
    input := make([]int, 0)
    for s.Scan() {
        t := s.Text()
        num, err := strconv.ParseInt(t, 0, 64)
        if err != nil {
            panic("non convertable int")
        }
        input = append(input, int(num))
    }
    var sum int = 2020
    res := ThreeSum(input, sum)
    if len(res) == 0 {
        return
    }
    fmt.Println(res[0] * res[1] * res[2])
}

func ThreeSum(nums []int, sum int) []int {
	sort.Ints(nums)
    for i := 0 ; i < len(nums) ; i++ {
        j := i + 1
        k := len(nums) - 1 
        for j < k {
            s := nums[i] + nums[j] + nums[k] 
            if s == sum {
                return []int{nums[i],nums[j], nums[k]}
            } else if s < sum {
                j++ 
            } else {
                k--
            }
        }
    }
    return []int{}
}
