package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    var s string
    s = solution("abcde")
    fmt.Fprintln(writer, s)
    
    s = solution("qwer")
    fmt.Fprintln(writer, s)
}

func solution(s string) string {
    var res string
    var size int = len(s)
    var isOdd bool = (size % 2) == 1
    size /= 2

    if isOdd {
        res = string(s[size])
    } else {
        res = string(s[size-1:size+1])
    }

    return res
}
