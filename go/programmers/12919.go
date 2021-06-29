package main

import (
    "fmt"
)

func main() {
    var a string
    a=solution([]string{"a234", "Kim"})
    fmt.Println(a)
    a=solution([]string{"a234", "Kim"})
    fmt.Println(a)
}

func solution(s []string) string {
    for i, v := range s {
        if v != "Kim" {
            continue
        }
        return fmt.Sprintf("김서방은 %d에 있다", i)
    }

    return ""
}
