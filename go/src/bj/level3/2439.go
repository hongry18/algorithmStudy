package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

var (
    w = bufio.NewWriter(os.Stdout)
    r = bufio.NewReader(os.Stdin)
)

func main() {
    defer w.Flush()
    var n int
    fmt.Fscanf(r, "%d\n", &n)

    var strFormat string = fmt.Sprintf("%s%ds\n", "%", n)

    for i:=0; i<n; i++ {
        fmt.Fprintf(w, strFormat, strings.Repeat("*", (i+1)))
    }
}
