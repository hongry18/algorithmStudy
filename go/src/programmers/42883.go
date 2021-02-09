package main

import (
    "fmt"
    "strconv"
)

func main() {
    var a string
    a=solution("1924", 0)
    //a=solution("87654321", 3)
    //a=solution("4177252841", 4)
    //a=solution("1231234", 3)
    //a=solution("01010", 3)
    //a=solution("12345678901234567890", 12)
    //a=solution("0102", 3)
    fmt.Println(a)
}

func solution(s string, n int) string {
    if n == 0 {
        return s
    }

    var length int = len(s)

    if length < 2 {
        return s
    }

    if length == n {
        return s
    }

    var c rune = 47
    var size int = length - n
    var idx int = 0
    var ll int = 0
    var f int = 0
    var e = size
    var res []byte

    if size == 1 {
        for i, v := range s {

            if v > c {
                c = v
                f = i
            }
        }

        return string(s[f])
    }

    if size < n {
        for i, v := range s[0: n] {
            if v > c {
                c = v
                f = i
            }
        }

        s = s[f:]
        length = len(s)
        f = 0
    }

    c = 47

    if f > e {
        return s[f:]
    }

    for {
        if e > length {
            res = append(res, s[f:f+size]...)
            break
        }

        for i, v := range s[f: e] {
            if v == 57 {
                idx = f + 1
                break
            }

            if v > c {
                c = v
                idx = f + i
            }
        }

        res = append(res, s[idx])

        c = 47
        size--
        f = idx + 1
        ll = len(s[f:])
        e = idx + ll - size + 2

        if e < f {
            res = append(res, s[f:]...)
            break
        }

        if size == ll {
            res = append(res, s[f:]...)
            break
        }
    }

    if res[0] == 48 {
        dd, _ := strconv.Atoi(string(res))
        return strconv.Itoa(dd)
    }

    return string(res)
}
