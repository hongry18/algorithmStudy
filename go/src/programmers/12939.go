package main

import "fmt"
import "strings"
import "strconv"

func main() {
    var a string
    a = solution("1 2 3 -4")
    fmt.Println(a)
}

func solution(s string) string {

    ss := strings.Split(s, " ")
    min, max := 999999999999999, -999999999999999
    for _, v := range ss {
        c, _ := strconv.Atoi(v)
        if c > max {
            max = c
        }

        if c < min {
            min = c
        }
    }


    return fmt.Sprintf("%d %d", min, max)
}
