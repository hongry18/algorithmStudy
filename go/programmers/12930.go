package main

import "fmt"
import "strings"
import "bytes"

func main() {
    var a string
    a=solution("try hello world")
    fmt.Println(a)
}

func solution(s string) string {
    var a []string = strings.Split(s, " ")
    var buf bytes.Buffer
    var size int = len(a) - 1
    for i, v := range a {
        var b []byte = []byte(v)
        for bi, bv := range b {
            b[bi] = swapCase(bv, bi % 2 == 1)
        }

        buf.WriteString(string(b))
        if size != i {
            buf.WriteString(" ")
        }
    }

    return buf.String()
}

func swapCase(a byte, f bool) byte {
    var diff int = 32
    /*
97 a
122 z

65 A
90 Z
    */

    // to lower case
    if f && a < 97{
        return byte(int(a) + diff)
    }

    // to upper case
    if !f && a > 90{
        return byte(int(a) - diff)
    }

    return a
}
