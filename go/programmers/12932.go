package main

import "fmt"

func main() {
    var a []int
    a=solution(123)
    fmt.Println(a)
    a=solution(987)
    fmt.Println(a)
    a=solution(102004000)
    fmt.Println(a)
    a=solution(999999999)
    fmt.Println(a)
    a=solution(10208050)
    fmt.Println(a)
    a=solution(987654321)
    fmt.Println(a)
}

func solution(n int64) []int {
    var size int = Size(n)
    var res = make([]int, size)

    for i:=0; i<size; i++ {
        idx := size - i
        c:=Pow(idx)
        a := n/c
        n%=c
        res[idx-1] = int(a)
    }

    return res
}

func Size(n int64) (size int) {
    size = 0
    var start int64 = 1

    for {
        if n < start {
            return size
        }

        start *= 10
        size++
    }

    return size
}

func Pow(n int) (a int64) {
    a = 1
    var end int = n-1

    for i:=0; i<end; i++ {
        a *= 10
    }

    return a
}
