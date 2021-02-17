package main

import "fmt"

func main() {
    var a [][]int
    //a = solution([][]int{{1, 4}, {3, 2}, {4, 1}}, [][]int{{3, 3}, {3, 3}})
    a = solution([][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}})
    fmt.Println("Answer: ", a)
}


// 행렬의 곱셈은 좀 외워두자
func solution(a1, a2 [][]int) [][]int {

    var res [][]int = make([][]int, len(a1))

    for i:=0; i<len(a1); i++ {
        res[i] = make([]int, len(a2[0]))
        for j:=0; j<len(res[0]); j++ {
            for k:=0; k<len(a1[0]); k++ {
                res[i][j] += a1[i][k] * a2[k][j]
            }
        }
    }

    return res
}

