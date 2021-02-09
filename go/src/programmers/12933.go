package main

import "fmt"
import "sort"

func main() {
    var a int64
    a=solution(118372)
    fmt.Println(a)
    a=solution(987)
    fmt.Println(a)
/*
    a=solution(102004000)
    fmt.Println(a)
    a=solution(999999999)
    fmt.Println(a)
    a=solution(10208050)
    fmt.Println(a)
    a=solution(987654321)
    fmt.Println(a)
*/
}

func solution(n int64) int64 {
    var size int = Size(n)
    var res []int
    var sum int64 = 0

    for i:=0; i<size; i++ {
        idx := size - i
        c:=Pow(idx)
        a := n/c
        n%=c
        res = append(res, int(a))
    }

    sort.Slice(res, func(x, y int) bool {return res[x] < res[y]})
    for i, v := range res {
        sum += int64(v) * Pow(i+1)
    }

    return sum
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
