package main

import "fmt"

func main() {
    var a int
    a = solution(30000)

    fmt.Println(a)
}

func solution(n int) int {
    a := fibo(n)
    return int(a % 1234567)
}

func fibo(n int) int {

    if n < 2 {
        return n
    }

    var a, b, c int = 0, 1, 1

    for i:=0; i<n-2; i++ {
        a = b % 1234567
        b = c % 1234567
        c = a + b
    }

    return c
}
