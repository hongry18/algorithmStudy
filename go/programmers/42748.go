package main

import (
    "fmt"
    "sort"
)

func main() {
    var s []int
    s = solution([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2,5,3}, {4,4,1}, {1,7,3}})
    fmt.Println(s)
}

func solution(array []int, commands [][]int) []int {
    res := []int{}
    size := len(commands)
    for i:=0; i<size; i++ {
        f := make([]int, len(array))
        copy(f,array)
        n := f[commands[i][0]-1:commands[i][1]]
        sort.Ints(n)
        res = append(res, n[commands[i][2]-1])
    }

    return res
}
