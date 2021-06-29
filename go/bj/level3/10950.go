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

    var n, a, b int
    fmt.Fscanf(r, "%d\n", &n)

    for i:=0; i<n; i++ {
        fmt.Fscanf(r, "%d %d\n", &a, &b)

        fmt.Fprintf(w, "%d\n", (a+b))
    }
}
