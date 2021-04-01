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

    var n, a, b int

    fmt.Fscanf(r, "%d\n", &n)

    for i:=0; i<n; i++ {
        fmt.Fscanf(r, "%d %d\n", &a, &b)
        fmt.Fprintf(w, "Case #%d: %d + %d = %d\n", (i + 1), a, b, (a + b))
    }
}
