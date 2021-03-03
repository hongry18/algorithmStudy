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
    fmt.Fscanf(r, "%d\n%d\n", &a, &b)

    if a > 0 && b > 0 {
        fmt.Fprintln(w, 1)
        return
    }

    if a < 0 && b > 0 {
        fmt.Fprintln(w, 2)
        return
    }

    if a < 0 && b < 0 {
        fmt.Fprintln(w, 3)
        return
    }

    fmt.Fprintln(w, 4)
}
