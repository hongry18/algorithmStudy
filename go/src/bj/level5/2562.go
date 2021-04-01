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

    var n, max, idx, maxIdx int
    max = -100
    idx = 1

    for {
        _, err := fmt.Fscanf(r, "%d\n", &n)

        if n > max {
            max = n
            maxIdx = idx
        }

        idx++
        if err != nil {
            break
        }
    }

    fmt.Fprintf(w, "%d\n", max)
    fmt.Fprintf(w, "%d\n", maxIdx)
}
