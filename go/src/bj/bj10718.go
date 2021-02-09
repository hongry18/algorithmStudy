package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    var t string = "강한친구 대한육군"
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    for i:=0; i<2; i++ {
        fmt.Fprintf(writer, "%s\n", t)
    }
}
