package main

import (
    "fmt"
    "bytes"
)

func main() {
    var a string
    a=solution(0)
    fmt.Println(a)
    a=solution(1)
    fmt.Println(a)
    a=solution(3)
    fmt.Println(a)
    a=solution(4)
    fmt.Println(a)
    a=solution(9)
    fmt.Println(a)
}

func solution(n int) string {
    var l int = n / 2
    var isOdd bool = n % 2 == 1

    var b bytes.Buffer

    for i:=0; i<l; i++ {
        b.WriteString("수박")
    }

    if isOdd {
        b.WriteString("수")
    }

    return b.String()
}
