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

    var n int
    fmt.Fscanf(r, "%d\n", &n)

    for i:=0; i<n; i++ {
        fmt.Fprintf(w, "%d\n", (i+1))
    }
}
