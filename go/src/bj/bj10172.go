package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    txt := "|\\_/|\n"
    txt += "|q p|   /}\n"
    txt += "( 0 )\"\"\"\\\n"
    txt += "|\"^\"`    |\n"
    txt += "||_/=\\\\__|\n"

    fmt.Fprintf(writer, "%s", txt)
}
