package main

import "fmt"

func main() {
    var s int64
    s = solution(3,5)
    fmt.Println(s)
    s = solution(3,3)
    fmt.Println(s)
    s = solution(5,3)
    fmt.Println(s)
}

func solution(a, b int) int64 {
    aa := int64(a)
    bb := int64(b)
    if aa == bb {
        return aa
    }

    if bb > aa {
        return ((aa + bb) * ((bb - aa) + 1)) / 2
    }

    return ((aa + bb) * ((aa - bb) + 1)) / 2
}
