package main

import (
    "fmt"
)

func main() {
    var r int
    r = solution(5, []int{2, 4}, []int{1,3,5})
    fmt.Println(r)
    r = solution(5, []int{2, 4}, []int{3})
    fmt.Println(r)
    r = solution(3, []int{3}, []int{1})
    fmt.Println(r)
    r = solution(2, []int{1,2}, []int{1,2})
    fmt.Println(r)
    r = solution(3, []int{1,2}, []int{2,3})
    fmt.Println(r)
}

func solution(n int, lost []int, reserve []int) int {
    var uok, rok bool
    used := make(map[int]bool, n)
    r2 := make(map[int]bool, n)
    var res []int

    size := len(reserve)
    for i:=0; i<size; i++ {
        r2[reserve[i]] = true
    }

    size = len(lost)
    maxSize := n-1

    for i:=0; i<size; i++ {
        idx := lost[i]
        _, rok = r2[idx]
        if rok {
            used[idx] = true
        }
    }

    for i:=0; i<size; i++ {
        idx := lost[i]
        _, rok = r2[idx]
        if rok {
            continue
        }

        // left
        if idx > 1 {
            _, uok = used[idx-1]
            _, rok = r2[idx-1]
            if !uok && rok {
                used[idx-1] = true
                continue
            }
        }

        // right
        if idx <= maxSize {
            _, uok = used[idx+1]
            _, rok = r2[idx+1]
            if !uok && rok {
                used[idx+1] = true
                continue
            }
        }

        res = append(res, idx)
    }

    return n - len(res)
}
