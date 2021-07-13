package main

import(
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

    if a > b {
        fmt.Fprintf(w, ">")
        return
    }

    if a < b {
        fmt.Fprintf(w, "<")
        return
    }

    fmt.Fprintf(w, "==")
}
