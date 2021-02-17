package main

import "fmt"

func main() {
    var a int

    a = solution([]int{1,1,1,1,1}, 3)
    fmt.Println(a)
}

func solution(ar []int, t int) int {
    var cnt int = 0

    var a []int = []int{0}
    for i:=0; i<len(ar); i++ {
        fmt.Println(a, ar[i])
        a = get(a, ar[i])
    }

    for _,v := range a {
        if v != t {
            continue
        }
        cnt++
    }

    return cnt
}

func get(a []int, p int) (res []int) {

    for _,v := range a {
        res = append(res, v + p)
        res = append(res, v + (-1 * p))
    }

    return
}
