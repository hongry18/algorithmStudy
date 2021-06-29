package main

import "fmt"

func main() {
    var a int
    a=solution(12)
    fmt.Println(a)

    a=solution(5)
    fmt.Println(a)
}

func solution(n int) int {
    var sum int = n
    for i:=1; i<n; i++ {
        if n % i != 0 {
            continue
        }
        sum += i
    }
    return sum
}

