package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    var a int
    var b int

    r := bufio.NewReader(os.Stdin)
    fmt.Fscanf(r, "%d %d", &a, &b)

    fmt.Print(a+b)
}
