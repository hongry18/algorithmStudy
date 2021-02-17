package main

import "fmt"

func main() {
    var a int
    //a=solution([]int{3,0,6,1,5})
    a=solution([]int{5,5,5,5})

    fmt.Println("Answer: ", a)
}

func solution(c []int) int {
    var length int = len(c)
    var max int = 0
    var idx int = 0

    for i:=length; i>-1; i-- {
        for _, v := range c {
            if v < i {
                continue
            }
            max++
        }

        if max >= i {
            idx = i
            break
        }

        max = 0
    }

    return idx
}
