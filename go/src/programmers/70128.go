package main

import "fmt"

func main() {
    var a int
    a=solution([]int{1,2,3,4}, []int{-3, -1, 0, 2})
    fmt.Println(a)
}

func solution(a, b []int) int {
    var sum int = 0
    for i := range a {
        sum += a[i]*b[i]
    }

    return sum
}
