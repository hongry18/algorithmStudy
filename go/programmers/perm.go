package main

import (
    "fmt"
)

func Permutations(L []rune, r int) [][]rune {
    if r == 1 {
        //Convert every item in L to List and
        //Append it to List of List
        temp := make([][]rune, 0)
        for _, rr := range L {
            t := make([]rune, 0)
            t = append(t, rr)
            temp = append(temp, [][]rune{t}...)
        }
        return temp
    } else {
        res := make([][]rune, 0)
        for i := 0; i < len(L); i++ {
            //Create List Without L[i] element
            perms := make([]rune, 0)
            perms = append(perms, L[:i]...)
            perms = append(perms, L[i+1:]...)
            //Call recursively to Permutations
            for _, x := range Permutations(perms, r-1) {
                t:=append(x, L[i])
                res=append(res, [][]rune{t}...)
            }
        }
        return res
    }
}

func main() {
    res := Permutations([]rune("1234"), 2)
    fmt.Println(res, len(res))

}
