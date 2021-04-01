package main

import (
    "os"
    "fmt"
    "bufio"
)

var (
    reader = bufio.NewReader(os.Stdin)
    writer = bufio.NewWriter(os.Stdout)
)

func main() {
    defer writer.Flush()
    var n int;
    fmt.Fscanf(reader, "%d\n", &n)

    for i:=1; i<10; i++ {
        fmt.Fprintf(writer, "%d * %d = %d\n", n, i, n*i)
    }
}
