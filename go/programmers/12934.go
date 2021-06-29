package main

import (
    "fmt"
    "math"
)

func main() {
    solution(121)
    solution(3)
}

func solution(n int64) int64 {
    var a float64 = math.Sqrt(float64(n))
    var b float64 = float64(int64(a))
    if a != b {
        return -1
    }
    var c int64 = int64(a) + 1
    return c * c
}
