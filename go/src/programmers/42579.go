package main

import "fmt"
import "sort"

func main() {
    var a []int
    //a = solution([]string{"classic", "pop", "classic", "classic", "pop"}, []int{500, 600, 150, 800, 2500})
    //a = solution([]string{"classic", "classic", "classic", "classic", "classic"}, []int{500, 600, 150, 800, 2500})
    //a = solution([]string{"classic"}, []int{500})
    a = solution([]string{"classic", "pop", "classic", "pop", "classic", "classic"}, []int{400, 600, 150, 600, 500, 500})
    fmt.Println(a)
}


func solution(g []string, p []int) []int {

    items := make(map[string][]int)
    sums := make(map[string]int)
    names := make([]string, 0)
    nums := make(map[int]int)
    res := make([]int, 0)

    for i, v := range g {
        item, ok := items[v]; if !ok {
            item = make([]int, 0)
        }
        sum, ok := sums[v]; if !ok {
            sum = 0
        }
        num, ok := nums[p[i]]; if !ok {
            num = i
            nums[p[i]] = num
        }

        item = append(item, p[i])
        items[v] = item
        sum += p[i]
        sums[v] = sum
    }
    fmt.Println(nums)

    for i := range sums {
        names = append(names, i)
    }

    sort.Slice(names, func(a, b int) bool {
        return sums[names[a]] > sums[names[b]]
    })

    for i := range names {
        a := items[names[i]]
        if len(a) == 1 {
            res = append(res, nums[a[0]])
            continue
        }
        sort.Sort(sort.Reverse(sort.IntSlice(a)))
        res = append(res, nums[a[0]])
        res = append(res, nums[a[1]])
    }

    return res
}
