package main

import (
    "fmt"
    "sort"
)

func main() {
    var s []int
    s = solution([]int{5,9,7,10}, 5)
    fmt.Println(s)
    s = solution([]int{2, 36, 1, 3}, 1)
    fmt.Println(s)
    s = solution([]int{3, 2, 6}, 10)
    fmt.Println(s)
}

func solution(arr []int, divisor int) []int {
    var res []int
    for _, i := range arr {
        if i % divisor != 0 {
            continue
        }

        res = append(res, i)
    }

    if len(res) < 1 {
        res = append(res, -1)
    } else {
        sort.Ints(res)
    }

    return res
}
