package main

import (
    "os"
    "fmt"
    "bufio"
)

var (
    w = bufio.NewWriter(os.Stdout)
    r = bufio.NewReader(os.Stdin)
)

func main() {
    defer w.Flush()
    var n, s, a, b, c, o int
    var cnt int = 1

    fmt.Fscanf(r, "%d\n", &n)
    o = n

    if n == 0 {
        fmt.Fprintf(w, "%d\n", 1)
        return
    }

    if n > 99 {
        return
    }

    for {
        a, b = div(n)
        s = a + b
        if s > 9 {
            _, c = div(s)
            n = b*10 + c
        } else {
            n = b*10 + s
        }

        //fmt.Fprintf(w, "cnt: %d, %d - %d\n", n, o, cnt)
        if n == o {
            break
        }
        cnt++
    }

    fmt.Fprintf(w, "%d\n", cnt)
}

func div(x int) (a, b int) {
    if x < 10 {
        a = 0
        b = x
        return
    }

    a = x / 10
    b = x % 10
    return
}
