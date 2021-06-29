package main

import "fmt"

func main() {
    var a []int

    //a = solution("110010101001")
    //a = solution("1111111")
    a = solution("01110")
    fmt.Println(a)
}

func solution(s string) []int {

    cnt := 0
    loop := 0
    b := s

    for {
        c, a := removeZero([]byte(b))
//        fmt.Printf("c: %d, a: %s, b: %s\n", c, a, b)
        cnt += c

        b = genB([]byte(a))
        loop++

        if len(a) < 1 {
            break
        }

        if a == "1" {
            break
        }
    }

    return []int{loop, cnt}
}

func removeZero(a []byte) (int, string) {
    var cnt int = 0
    var res []byte
    for _, v := range a {
        if v == 48 {
            cnt++
            continue
        }

        res = append(res, v)
    }

    return cnt, string(res)
}

func genB(a []byte) string {
    var c int = len(a)
    var res []byte

    for {
        if c < 2 {
            res = append(res, byte(c+48))
            break
        }
        b := c % 2
        res = append(res, byte(b+48))
        c /= 2
    }

    return string(res)
}
