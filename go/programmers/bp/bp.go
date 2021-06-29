package bp

import (
    "os"
    "fmt"
    "bufio"
)

func Print(f string, i ...interface{}) {
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    fmt.Fprintf(writer, f, i...)
}
