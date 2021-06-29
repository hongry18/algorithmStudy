package main

import (
    "fmt"
    "sort"
)

func main() {
    var s []string
    s = solution([]string{"sun", "bed", "car"}, 1)
    fmt.Println(s)
    s = solution([]string{"abce", "abcd", "cdx"}, 2)
    fmt.Println(s)
}

func solution(strs []string, n int) []string {
    var res []string
    var keys []string
    var items = make(map[string][]string)
    for _, i := range strs {
        var key string = string(i[n])
        items[key] = append(items[key], i)
    }

    for k, _ := range items {
        keys = append(keys, k)
    }

    sort.Strings(keys)

    for _, k := range keys {
        if len(items[k]) > 1 {
            sort.Strings(items[k])
        }
        res = append(res, items[k]...)
    }

    return res
}
