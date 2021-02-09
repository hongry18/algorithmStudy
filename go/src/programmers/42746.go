package main

import "fmt"
import "strconv"
import "sort"
import "strings"
import "bytes"

func main() {
//    solution([]int{21, 212})
//    solution([]int{0, 0, 1000})
//    solution([]int{6, 10, 2})
//    solution([]int{3, 30, 34, 5, 9})
//    solution([]int{12, 121})
//    solution([]int{21, 212})

}

func solution(arr []int) string {
    var sum int = 0
    var sarr []string
    for _, v := range arr {
        sum += v
        sarr = append(sarr, strconv.Itoa(v))
    }

    if sum == 0 {
        return "0"
    }

    sort.Slice(sarr, func(x, y int) bool {
        if (sarr[x][0] > sarr[y][0]) {
            return true
        }

        return sarr[x] + sarr[y] > sarr[y] + sarr[x]
    })

    fmt.Println(sarr)

    return strings.Join(sarr, "")
}
