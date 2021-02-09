package main

import (
    "fmt"
)

func main() {
    solution("a234")
    solution("0122030z349")
}

func solution(s string) bool {
    size := len(s)

    if size != 4 && size != 6 {
        return false
    }

    for _, i := range s {
        if i < '0' || i > '9' {
            return false
        }
    }

    return true
}
