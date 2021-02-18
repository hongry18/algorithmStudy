package main

import "fmt"
import "math"

func main() {
    var a int
    //a = solution(3, [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
    //a = solution(3, [][]int{{1, 1, 0}, {1, 1, 1}, {0, 0, 1}})
    a = solution(3, [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})
    //a = solution(3, [][]int{{1, 0, 1, 1, 0, 0}, {0, 1, 0, 0, 1, 1},{0, 1, 0, 0, 1, 1}, {1, 0, 1, 1, 1, 1}, {1, 0, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1}})
    fmt.Println(a)

}

func solution(n int, c [][]int) int {
    return find(c[1:], ToInt(c[0]))
}

func find(c [][]int, a int) int {
    var o int = a
    var b int
    var unmatch [][]int
    var sum = 1

    for i:=0; i<len(c); i++ {
        b = ToInt(c[i])
        if a & b == 0 {
            unmatch = append(unmatch, c[i])
            continue
        }

        a = a | b
    }

    if len(unmatch) == 0 {
        return sum
    }

    d := ToInt(unmatch[0])
    if len(unmatch) == 1 {
        if a & d == 0 {
            return sum + 1
        } else {
            return sum
        }
    }

    if a & d == 0 {
        return find(unmatch[1:], d)
    }

    if o == a {
        sum--
    }

    return find(unmatch[1:], d|a)
}

func ToInt(a []int) int {
    var length int = len(a)
    var sum float64 = 0

    for i:=0; i<length; i++ {
        if a[i] == 0 {
            continue
        }
        sum += math.Pow(2, float64(length - i - 1))
    }
    return int(sum)
}
