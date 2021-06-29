package main

import (
    "fmt"
)

func main() {
    solution([]int{2,1,3,2}, 2)
    solution([]int{1,1,9,1,1,1}, 0)
}

type Queue struct {
    i [][]int
}

func (q *Queue) add(d []int) {
    q.i = append(q.i, d)
}

func (q *Queue) remove() (res []int) {
    if q.size() < 1 {
        return []int{}
    }
    res = q.i[0]
    q.i = q.i[1:]
    return
}

func (q *Queue) size() int {
    return len(q.i)
}

func (q *Queue) peek() []int {
    if q.size() < 1 {
        return []int{}
    }
    return q.i[0]
}

func solution(ps []int, loc int) (res int) {
    res = 1
    q := Queue{i: make([][]int, len(ps))}
    isBig := false

    for k,v := range ps {
        q.i[k] = []int{v, k}
    }

    for {
        if q.size() < 2 {
            break
        }

        a := q.remove()
        isBig = false
        for i:=0; i<q.size(); i++ {
            if q.i[i][0] > a[0] {
                isBig = true
                break
            }
        }

        if isBig {
            q.add(a)
            continue
        }

        if a[1] == loc {
            break
        }
        res++
    }

    fmt.Println(res)

    return
}
