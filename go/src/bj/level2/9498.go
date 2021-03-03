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
    var a int
    fmt.Fscanf(r, "%d\n", &a)

    if a > 89 {
        fmt.Fprintln(w, "A")
        return
    }

    if a > 79 {
        fmt.Fprintln(w, "B")
        return
    }

    if a > 69 {
        fmt.Fprintln(w, "C")
        return
    }

    if a > 59 {
        fmt.Fprintln(w, "D")
        return
    }

    fmt.Fprintln(w, "F")
}
