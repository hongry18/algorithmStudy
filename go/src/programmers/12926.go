package main

import "fmt"

func main() {
    var a string
    a=solution("AB", 1)
    fmt.Println(a)

    a=solution("z", 1)
    fmt.Println(a)

    a=solution("a B z", 4)
    fmt.Println(a)

}

func solution(s string, n int) string {
    var b []byte = []byte(s)

    // 122 = z 97 = a
    // 90 = Z 65 = A

    for i, v := range b {
        if v == 32 {
            continue
        }

        c := int(b[i]) + n

        if (b[i] < 91 && c > 90) || (b[i] > 96 && c > 122) {
            c -= 26
        }

        b[i] = byte(c)
    }
    return string(b)
}
