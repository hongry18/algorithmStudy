package main

import "fmt"

import "strings"

func main() {
    solution("01033334444")
    solution("027778888")
}

func solution(s string) string {
    var size int = len(s)
    var a []string = []string{}
    for i:=0; i<size-4; i++ {
        a = append(a, "*")
    }
    a = append(a, s[size-4: size])
    return strings.Join(a, "")
}
