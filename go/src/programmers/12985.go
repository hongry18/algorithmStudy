package main

import "fmt"

func main() {
    var a int
    //a = solution(8, 1, 2)
    //a = solution(8, 2, 3)
    a = solution(4096, 2, 3999)
    //a = solution(32, 8, 10)
    fmt.Println(a)
    a = solution3(4096, 2, 3999)
    fmt.Println(a)
}

func solution(n, a, b int) int {
    var cnt int = 1

    for {
        if a % 2 == 0 && a - 1 == b {
            break
        }

        if a % 2 != 0 && a + 1 == b {
            break
        }

        a = ton(a)
        b = ton(b)

        cnt++
    }

    return cnt
}

func ton(x int) int {
    if x % 2 == 0 {
        return x / 2
    }
    return (x / 2) + 1
}

func solution3(n, a, b int) int {
    if n == 2 {
        return 1
    }

    if a < b {
        a, b = b, a
    }

    if a % 2 == 0 && a - b == 1 {
        return 1
    }

    if a % 2 == 1 && a - b == 1 {
        return 2
    }

    var cnt int = 1
    var i int = 2
    var c int = a - b

    if a % 2 == 0 {
        c += 1
    } else {
        c += 2
    }

    for {
        if i >= c {
            break
        }
        i *= 2
        cnt++
    }

    return cnt
}

func solution2(n, a, b int) int {
    var ar []int = make([]int, n)
    if n == 2 {
        return 1
    }

    var stop bool = false
    var length int = n / 2
    var cnt int = 0

    for i:=0; i<n; i++ {
        ar[i] = i+1
    }

    for length != 0 && !stop {
        for i:=length; i>0; i-- {
            var idx int = i*2-2

            if ar[idx] == a && ar[idx+1] == b {
                stop=true
                break
            }

            if ar[idx] == b && ar[idx+1] == a {
                stop=true
                break
            }

            if ar[idx+1] == a || ar[idx+1] == b {
                ar = append(ar[:idx], ar[idx+1:]...)
                continue
            }

            ar = append(ar[:idx+1], ar[idx+2:]...)
        }
        cnt++
        length /= 2
    }

    return cnt
}
