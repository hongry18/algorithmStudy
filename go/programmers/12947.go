package main

import "fmt"

func main() {
    solution(10)
    solution(12)
    solution(11)
    solution(13)
}

func solution(n int) bool {
    var ori int = n
    var size int = Size(n)
    var d int = getDiv(size)
    var sum int = 0

    for i:=0; i<size; i++ {
        sum += n / d
        n = n % d
        d = d / 10
    }

    return ori % sum == 0
}

func Size(n int) int {
    var i int = 10
    var cnt int = 1
    for {
        if n < i {
            break
        }
        i *= 10
        cnt ++
    }
    return cnt
}

func getDiv(n int) int {
    var r int = 1
    for i:=1; i<n; i++ {
        r *= 10
    }

    return r
}
