package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    reader := bufio.NewReader(os.Stdin)

    var a int
    var b int

    fmt.Fscan(reader, &a, &b)

    fmt.Fprintf(writer, "%d\n", (a+b))
    fmt.Fprintf(writer, "%d\n", (a-b))
    fmt.Fprintf(writer, "%d\n", (a*b))
    fmt.Fprintf(writer, "%d\n", (a/b))
    fmt.Fprintf(writer, "%d\n", (a%b))
}
