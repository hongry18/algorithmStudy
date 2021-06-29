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
    for i:=0; i<2; i++ {
        fmt.Fprintf(w, "%s\n", "강한친구 대한육군")
    }
}
