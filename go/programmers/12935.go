package main

import "fmt"

func main() {
    solution([]int{4,3,2,1})
    solution([]int{10, 10, 10})
}

func solution(a []int) []int {
    var size int = len(a)
    if size == 1 {
        return []int{-1}
    }

    var min = a[0]
    var res []int = []int{}

    for i:=1; i<size; i++ {
        if a[i] < min {
            min = a[i]
        }
    }

    for _, v := range a {
        if v == min {
            continue
        }
        res = append(res, v)
    }

    if len(res) < 1 {
        return []int{-1}
    }

    return res
}
