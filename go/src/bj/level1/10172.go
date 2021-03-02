package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    txt := `|\_/|
|q p|   /}
( 0 )"""\
|"^"` + "`" + `    |
||_/=\\__|
`

    fmt.Fprintf(writer, "%s", txt)
}
