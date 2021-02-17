package main

import "fmt"

func main() {
    var a string
    //a = solution("3people unFollowed me")
    //a = solution("aaaaa aaa")
    //a = solution("a  3Sdf ffTt A fftt    fftt  3133Sdf    sd85f1    ")
    a = solution("a  B3Sdf ZffTt A fftt  233333  D z fftt  3133Sdf    sd85f1    ")
    fmt.Println("Answer: [", a, "]")
}

func solution(s string) string {
    var b []byte = []byte(s)

    for i, v := range b {
        if i == 0 && v > 96 && v < 123 {
            b[i] -= 32
            continue
        }

        if i != 0 && b[i-1] == 32 && v > 96 && v < 123 {
            b[i] -= 32
            continue
        }

        if i != 0 && b[i-1] != 32 && v > 64 && v < 91 {
            b[i] += 32
        }
        
    }

    return string(b)
}
