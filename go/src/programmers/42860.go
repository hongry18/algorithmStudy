package main

import "fmt"

func main() {
    //solution("JEROEN")
    solution("JAN")
}

func solution(s string) int {
    var (
        prev int = 0
        left int = 0
        right int = 0

        items []byte = []byte(s)

        size int = len(items)

        c int
    )

    c = calc(items[0])
    left = c
    right = c

    for i:=size-1; i>0; i-- {
        c = calc(items[i])
        if prev == 0 {
            prev = c
            continue
        }
        right += 1
        right += c
    }

    for i:=1; i<size; i++ {
        c = calc(items[i])
        if prev == 0 {
            prev = c
            continue
        }
        left += 1
        left += c
    }

    fmt.Println(left, right)
    return 0
}

func calc(c byte) int {
    r := int(c) - 65
    if r > 13 {
        r = 25 - r
    }
    return r
}
