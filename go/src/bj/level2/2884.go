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
    var a, b int
    fmt.Fscanf(r, "%d %d\n", &a, &b)

    if b > 44 {
        fmt.Fprintln(w, a, b-45)
        return
    }

    if a == 0 {
        a = 24
    }

    fmt.Fprintln(w, a-1, b+15)
}
