package main

import (
    "fmt"
)

func main() {
    var a string
    //a=solution("1924", 1)
    //a=solution("1231234", 3)
    //a=solution("87654321", 3)
    //a=solution("41797252841", 8)
    //a=solution("1231234", 3)
    //a=solution("01010", 3)
    //a=solution("12345678901234567890", 12)
    //a=solution("01020", 2)
    a=solution("00000000000000000000", 20)
    fmt.Println(a)
}

func solution(s string, n int) string {
    // 길이
    var sarr []byte = []byte(s)
    var length int = len(sarr)

    if length <= n {
        return ""
    }

    var find int = length - n
    var idx int = 0
    var end int = 0

    var res []byte = make([]byte, find)


    // 1 자리일때
    if length < 2 {
        return s
    }

    // 모든게 0일때
    if s[0] == 48 {
        for i:=0; i<length; i++ {
            if sarr[i] == 48 {
                continue
            }

            idx = i
            break
        }

        if idx == 0 {
            return "0"
        }
    }

    for i:=0; i<find; i++ {
        //fmt.Println("#####", i)
        end = length - ( find - i ) + 1
        //fmt.Printf("loop: [%d], idx: [%d], end: [%d], find: [%d], length: [%d], arr %s\n", i, idx, end, find-i, length, string(sarr[idx:end]))

        if end > length {
            end = length
        }


        if idx > end {
            idx = end - 1
        }

        idx = findMax(sarr, idx, end)
        //fmt.Printf("loop: [%d], idx: [%d], end: [%d], find: [%d], length: [%d], arr %s\n", i, idx, end, find-i, length, string(sarr[idx:end]))
        res[i] = sarr[idx]
        idx++

    }

    //fmt.Println(string(res))

    return string(res)
}

func findMax(a []byte, start, end int) int {
    var t byte = a[start]
    var idx int = start
    for i:=start; i<end; i++ {
        if a[i] == 57 {
            return i
        }

        if a[i] > t {
            t = a[i]
            idx = i
        }
    }

    return idx
}


/*
func solution2(s string, n int) string {
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
*/
