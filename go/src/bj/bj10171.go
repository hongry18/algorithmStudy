package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    txt := `\    /\
 )  ( ')
(  /  )
 \(__)|
`

    fmt.Fprintf(writer, "%s", txt)
}
