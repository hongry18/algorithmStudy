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

    var a, b int

    for {
        _, err := fmt.Fscanf(r, "%d %d\n", &a, &b)
        if err != nil {
            break
        }

        fmt.Fprintf(w, "%d\n", a+b)
    }
}
