package main

import (
    "fmt"
    "bytes"
)

func main() {
    fmt.Println(CombAll([]rune("abc")))
}

func CombAll(set []rune) (subsets []string) {
    length := uint(len(set))

    var buf bytes.Buffer
    for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
        for obj := uint(0); obj < length; obj++ {
            if (subsetBits >> obj)&1 == 1 {
                buf.WriteString(string(set[obj]))
            }
        }

        subsets = append(subsets, buf.String())
        buf.Reset()
    }

    return subsets
}
