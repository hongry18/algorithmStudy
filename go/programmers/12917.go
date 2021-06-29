package main

import (
    "fmt"
    "sort"
)

func main() {
    s:=solution("Zbcdefg")
    fmt.Println(s)
}

func solution(s string) string {
    a := []byte(s)
    sort.Slice(a, func(x, y int) bool {
        return a[x] > a[y]
    })

    return string(a)
}
