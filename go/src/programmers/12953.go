package main

import "fmt"

func main() {
    var a int
    a = solution([]int{2,6,8,14})
    a = solution([]int{1, 3, 4})
    a = solution([]int{4})
    fmt.Println("Answer: ", a)
}

func solution(a []int) int {

    if len(a) < 2 {
        return a[0]
    }

    return LCM(a[0], a[1], a[2:]...)
}

func GCD(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }

    return a
}

func LCM(a, b int, ints ...int) int {
    r := a * b / GCD(a, b)
    for i:=0; i<len(ints); i++ {
        r = LCM(r, ints[i])
    }
    return r
}
