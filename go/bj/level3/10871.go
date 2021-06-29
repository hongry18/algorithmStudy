package main

import (
    "os"
    "fmt"
    "bufio"
)

var (
    r = bufio.NewReader(os.Stdin)
    w = bufio.NewWriter(os.Stdout)
)

func main() {
    defer w.Flush()

    var n, x, y int
    fmt.Fscanf(r, "%d %d\n", &n, &x)

    for i:=0; i<n; i++ {
        if i == (n-1) {
            fmt.Fscanf(r, "%d", &y)
        } else {
            fmt.Fscanf(r, "%d ", &y)
        }

        if y >= x {
            continue
        }

        fmt.Fprintf(w, "%d ", y)
    }

    fmt.Fprint(w, "\n")
}
