package main

import "fmt"

func main() {
    //a = solution(8, 1, 2)
    //a = solution(8, 2, 3)
    //a = solution(32, 8, 10)

    var length = 8

    var items []int
    for i:=0; i<length; i++ {
        items = append(items, i+1)
    }
    var f [][]int = Comb(items, make([]bool, length), 0, length, 2)

    for _, v := range f {
        c1:=solution(length, v[0], v[1])
        c2:=solution3(length, v[0], v[1])
        if c1 == c2 {
            continue
        }
        fmt.Println(length, v[0], v[1], c1, c2)
    }
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
        if i > c {
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

func Comb(arr []int, visit []bool, s, n, r int) [][]int {
    if r == 0 {
        t := make([]int, 0)
        for i:=0; i<n; i++ {
            if visit[i] {
                t = append(t, arr[i])
            }
        }
        return [][]int{t}
    }

    res := make([][]int, 0)
    for i:=s; i<n; i++ {
        visit[i] = true
        c := Comb(arr, visit, i+1, n, r-1)
        res = append(res, c...)
        visit[i] = false
    }
    return res
}


func solution4(n int, a int, b int) int {
    ans := 1

    a--
    b--

    for a/2 != b/2 {
        ans++
        a /= 2
        b /= 2
    }

    return ans
}
