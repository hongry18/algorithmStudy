package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    fmt.Fprintf(writer, "%s\n", "Hello World!")
}
