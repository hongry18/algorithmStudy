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

    if a % 100 != 0 && a % 4 == 0 {
        fmt.Fprintln(w, 1)
        return
    }

    if a % 400 == 0 && a % 100 == 0 {
        fmt.Fprintln(w, 1)
        return
    }

    fmt.Fprintln(w, 0)
}
