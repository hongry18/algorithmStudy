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
    var n int
    var sum int64 = 0
    fmt.Fscanf(r, "%d\n", &n)

    for i:=1; i<=n; i++ {
        sum += int64(i)
    }
    fmt.Fprintf(w, "%d\n", sum)
}
