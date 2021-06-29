package main

import "fmt"
import "math"

func main() {
    var a []int
    //a=solution(5000, 4)
    //a=solution(24, 24)
    //a=solution(50, 22)
    a=solution(18, 6)
    //a=solution(16, 9)
    fmt.Println("answer: ", a)
}

func solution(b, y int) []int {
    var nb, ny int

    for i:=3; ; i++ {
        for j:=3; ; j++ {
            ny = (j-2) * (i-2)
            nb = 2*i + 2*j - 4

            if ny > y || nb > b {
                break
            }

            if ny == y && nb == b {
                if j > i {
                    i, j = j, i
                }
                return []int{i, j}
            }
        }
    }

    return []int{}
}

func solution1(b, y int) []int {

    var sum int = b + y
    var i int = 1

    for {
        if sum == ((y + (i * 2)) / i) * (i + 2) {
            break
        }
        i++
    }

    fmt.Println(i)
    ax := y / i + 2
    ay := i + 2

    if ay > ax {
        ax, ay = ay, ax
    }

    if (ax - 2) * (ay-2) == y {
        return []int{ax, ay}
    }

    return []int{}
}

func solution2(b, y int) []int {
    var sum int = b + y
    var f float64 = math.Sqrt(float64(sum))
    var i = int(f)
    var t int = 1
    var x int

    if f == float64(i) {
        return []int{i, i}
    }

    for {
        if i * (i+t) > sum {
            i--
            t=1
        }

        if i * (i+t) == sum {
            x = i+t
            break
        }

        t++
    }

    return []int{x, i}
}
