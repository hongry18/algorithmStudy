package main

import (
    "os"
    "fmt"
    "bufio"
)

var (
    w = bufio.NewWriter(os.Stdout)
)

func main() {
    defer w.Flush()
    fmt.Fprintf(w, "%s\n", "Hello World!")
}
