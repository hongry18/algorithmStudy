package main

import "fmt"
import "math"
import "errors"

func main() {
    var a int
    //a = solution(3, [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
    //a = solution(3, [][]int{{1, 1, 0}, {1, 1, 1}, {0, 0, 1}})
    a = solution(3, [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})
    //a = solution(3, [][]int{{1, 0, 1, 1, 0, 0}, {0, 1, 0, 0, 1, 1},{0, 1, 0, 0, 1, 1}, {1, 0, 1, 1, 1, 1}, {1, 0, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1}})
    fmt.Println(a)

}

func solution(n int, c [][]int) int {
    //return find(n, c)
    q := Queue{items: c}
    find(n, q)

    return 0
}

type Queue struct {
    items [][]int
}

func (q *Queue) Add(i []int) {
    q.items = append(q.items, i)
}

func (q *Queue) Get() ( []int, error ) {
    if len(q.items) == 0 {
        return []int{}, errors.New("empty")
    }

    item, items := q.items[0], q.items[1:]
    q.items = items
    return item, nil
}

func find(n int, q Queue) int {
    fmt.Println(q)
    var idx int = 0
    var item []int
    item, _ = q.Get()
    a := ToInt(item)
    for {
        if idx == n {
            break
        }

        item, e := q.Get()
        if e != nil {
            break
        }

        b := ToInt(item)

        if a & b == 0 {
            q.Add(item)
        }

        idx++
    }

    fmt.Println(a, idx)
    return 0
}

func ToInt(a []int) int {
    var length int = len(a)
    var sum float64 = 0

    for i:=0; i<length; i++ {
        if a[i] == 0 {
            continue
        }
        sum += math.Pow(2, float64(length - i - 1))
    }
    return int(sum)
}
