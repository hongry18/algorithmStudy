package main

import "fmt"
import "sort"

func main() {
    var a []int
    //a = solution([]string{"classic", "pop", "classic", "classic", "pop"}, []int{500, 600, 150, 800, 2500})
    //a = solution([]string{"classic", "classic", "classic", "classic", "classic"}, []int{500, 600, 150, 800, 2500})
    //a = solution([]string{"classic"}, []int{500})
    //a = solution([]string{"classic", "pop", "classic", "pop", "classic", "classic", "pop"}, []int{400, 600, 150, 600, 500, 500, 5000})
    //a = solution([]string{"classic", "pop", "classic", "pop", "classic", "classic"}, []int{400, 600, 150, 600, 500, 500})
    //a = solution([]string{"a", "b", "a", "b", "c"}, []int{100, 200, 300, 400, 500})
    a = solution([]string{"a", "b", "c", "d", "e", "f"}, []int{1, 2, 3, 4, 5, 6})
    fmt.Println(a)
}

type Item struct {
    nums map[int]int
    sum int
}

type kv struct {
    k int
    v int
}

func solution(g []string, p []int) []int {

    items := make(map[string]Item)
    nms := make([]string, 0)
    res := make([]int, 0)

    for i, v := range g {
        item, ok := items[v]; if !ok {
            item = Item{nums: make(map[int]int), sum: 0}
        }

        _, ok1 := item.nums[i]; if !ok1 {
            item.nums[i] = p[i]
        }
        item.sum += p[i]
        items[v] = item
    }

    for i := range items {
        nms = append(nms, i)
    }

    sort.Slice(nms, func(x, y int) bool {
        return items[nms[x]].sum > items[nms[y]].sum
    })

    for i := range nms {
        v := items[nms[i]]

        var ss []kv
        for k, v := range v.nums {
            ss = append(ss, kv{k, v})
        }

        sort.Slice(ss, func(x, y int) bool {
            if ss[x].v == ss[y].v {
                return ss[x].k < ss[y].k
            }
            return ss[x].v > ss[y].v
        })

        if len(ss) < 2{
            res = append(res, ss[0].k)
            continue
        }

        res = append(res, ss[0].k)
        res = append(res, ss[1].k)
    }


    return res
}
