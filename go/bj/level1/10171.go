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

    txt := `\    /\
 )  ( ')
(  /  )
 \(__)|
`

    fmt.Fprintf(w, "%s", txt)
}
