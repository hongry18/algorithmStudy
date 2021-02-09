package main

import (
    "fmt"
)

func main() {
    r:=Comb( []byte("abc"), make([]bool, 3), 0, 3, 2 )
    fmt.Println(r)
}

func Comb(arr []byte, visited []bool, s, n, r int) [][]byte {
    if r == 0 {
        t := make([]byte,0)
        for i:=0; i<n; i++ {
            if visited[i] {
                t = append(t, arr[i])
            }
        }
        res := [][]byte{t}
        
        return res
    }

    res := make([][]byte, 0)
    for i:=s; i<n; i++ {
        visited[i] = true
        c:=Comb(arr, visited, i+1, n, r-1)
        res = append(res, c...)
        visited[i] = false
    }

    return res
}
