package main

import (
    "os"
    "fmt"
    "bufio"
)

var cnts = []int{0,0}


func result(n int) (zeroCnt, oneCnt[]int) {
    zeroCnt = []int{1,0}
    oneCnt = []int{0,1}

    if n < 2 {
        return
    }

    for i:=2; i<n+1; i++ {
        zeroCnt = append(zeroCnt, zeroCnt[i-1] + zeroCnt[i-2])
        oneCnt = append(oneCnt, oneCnt[i-1] + oneCnt[i-2])
    }

    return
}

func main() {
    var t int
    reader := bufio.NewReader(os.Stdin)
    fmt.Fscanln(reader, &t)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    for i:=0; i<t; i++ {
        var n int
        fmt.Fscanln(reader, &n)
        zeroCnt, oneCnt := result(n)
        fmt.Fprintf(writer, "%d %d\n", zeroCnt[n], oneCnt[n])
    }
}

