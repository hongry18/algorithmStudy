package main

import "fmt"

func main() {
    solution(8, 12)
}

func solution(w, h int) int64 {
    w6 := int64(w)
    h6 := int64(h)

    answer := (w6 * h6) - (w6 + h6 - gcd(w6, h6))
    fmt.Println(answer)

    return answer
}

func gcd(a, b int64) int64 {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

func lcm(a, b int64) int64 {
    return a * b / gcd(a,b)
}
