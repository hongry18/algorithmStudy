package main

import "fmt"
import "strconv"

func main() {
    var a int
    a = solution("011")
    fmt.Println(a)
}

func solution(n string) int {

    var (
        cnt int
        exists bool
        items [][]rune
        dup map[int]bool = map[int]bool{0: false, 1: false}
    )

    for i := range n {
        items = append(items, Permutations([]rune(n), i+1)...)
    }

    for _, v := range items {
        c, _ := strconv.Atoi(string(v))

        _, exists = dup[c]

        if exists {
            continue
        }

        dup[c] = false
        if !IsPrime(c) {
            continue
        }

        cnt++
    }

    return cnt
}

func Permutations(list []rune, size int) (res [][]rune) {
    if size == 1 {
        for _, rr := range list {
            t := make([]rune, 0)
            t = append(t, rr)
            res = append(res, [][]rune{t}...)
        }
        return
    }

    for i:=0; i<len(list); i++ {
        perms := []rune{}
        perms = append(perms, list[:i]...)
        perms = append(perms, list[i+1:]...)

        for _,x := range Permutations(perms, size-1) {
            t := append(x, list[i])
            res = append(res, [][]rune{t}...)
        }
    }

    return
}

func IsPrime(n int) bool {
    if n < 2 {
        return false
    }

    if n < 4 {
        return true
    }

    if n % 2 == 0 || n % 3 == 0 {
        return false
    }

    for i:=3; i*i<=n; i+=2 {
        if n % i == 0 {
            return false
        }
    }

    return true
}
