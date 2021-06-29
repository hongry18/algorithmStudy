package main

import "fmt"

func main() {
    var a int
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

func solution(n int) int {
    var sum int = 0
    var size int = Size(n)
    for i:=0; i<size; i++ {
        c := cpow(size-i)
        a := n / c
        n = n % c
        sum += a
    }
    return sum
}

func Size(n int) int {
    var ck int = 1000000000
    var idx int = 10

    if n == ck {
        return idx
    }

    for i:=0; i<10; i++ {
        if n >= ck {
            return idx
        }
        idx--
        ck /= 10
    }

    return 0
}

func cpow(n int) int {
    var a int = 1
    for i:=0; i<n-1; i++ {
        a *= 10
    }

    return a
}
