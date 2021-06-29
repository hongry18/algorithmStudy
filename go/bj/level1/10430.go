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
    var a, b, c int
    fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)

    fmt.Fprintf(w, "%d\n", (a+b)%c)
    fmt.Fprintf(w, "%d\n", ((a%c)+(b%c))%c)
    fmt.Fprintf(w, "%d\n", (a*b)%c)
    fmt.Fprintf(w, "%d\n", ((a%c)*(b%c))%c)
}
