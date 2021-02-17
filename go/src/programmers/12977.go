package main

import "fmt"

func main() {
/*
    var a int
    a = solution([]int{1,2,7,6,4})
    fmt.Println(a)
*/

    fmt.Println("403" > "40")
}

func solution(ns []int) int {
    a := Comb(ns, make([]bool, len(ns)), 0, len(ns), 3)
    return len(a)
}

func Comb(arr []int, visit []bool, s, n, r int) []int {
    if r == 0 {
        t := 0
        for i:=0; i<n; i++ {
            if !visit[i] {
                continue
            }

            t += arr[i]
        }

        fmt.Println(t)
        if isPrime(t) {
            return []int{t}
        }

        return []int{}
    }

    res := make([]int, 0)
    for i:=s; i<n; i++ {
        visit[i] = true
        res = append(res, Comb(arr, visit, i+1, n, r-1)...)
        visit[i] = false
    }

    return res
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }

    if n < 4 {
        return true
    }

    if n % 2 == 0 || n % 3 == 0 {
        return false
    }

    for i:=3; i*i<=n; i+=2 {
        if n % i == 0 {
            return false
        }
    }

    return true
}
