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
    var a, b, c, d, e int
    c = 1
    e = 0
    fmt.Fscanf(r, "%d\n%d\n", &a, &b)

    for {
        if c > b {
            break
        }

        d = (b % (c * 10)) / c
        fmt.Fprintln(w, d*a)

        e += d * a * c

        c *= 10
    }

    fmt.Fprintln(w, e)

}
