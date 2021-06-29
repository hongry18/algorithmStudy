package main

import "fmt"
import "math"

func main() {
    var a int
    a=solution("1234")
    fmt.Println(a)
    a=solution("-1234")
    fmt.Println(a)
    a=solution("+1234")
    fmt.Println(a)
}

func solution(s string) int {
    var (
        size int = len(s)
        sum int = 0
        isPrefix bool = s[0] < 48
        isPos bool = s[0] != 45
        start int = 0
    )

    if isPrefix {
        start = 1
        size--
    }

    for i, v := range s[start:] {
        sum += int(math.Pow10(size - 1 - i)) * (int(v) - '0')
    }

    if isPos {
        return sum
    } else {
        return -1 * sum
    }
}
