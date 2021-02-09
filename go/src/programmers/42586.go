package main

import (
    "fmt"
)

type Queue []int
type Queue2 struct {
    items []int
}

func (q *Queue2) add(d int) {
    q.items = append(q.items, d)
}

func (q Queue2) peek() int {
    return q.items[0]
}

func (q *Queue2) pop() int {
    var p int = q.items[0]
    q.items = q.items[1:]
    return p
}

func (q *Queue2) size() int {
    return len(q.items)
}

func (q Queue) add(d int) Queue {
    return append(q, d)
}

func (q Queue) pop() (Queue, int) {
    return q[1:], q[0]
}

func (q Queue) peek() int {
    return q[0]
}

func main() {
    var a []int
/*
    a = solution([]int{93, 30, 55}, []int{1, 30, 5})
    fmt.Println(a)

    a = solution([]int{95, 30, 100, 55}, []int{1, 10, 30, 5})
    fmt.Println(a)

    a = solution([]int{93, 100, 55}, []int{1, 30, 5})
    fmt.Println(a)

    a = solution2([]int{93, 30, 55}, []int{1, 30, 5})
    fmt.Println(a)
*/
    var q = &Queue{}
    fmt.Println(q)

    a = solution2([]int{100, 93, 30, 55}, []int{40, 1, 30, 5})
    fmt.Println(a)
}

func solution2(p, s []int) []int {
    var res []int
    for i := range p {
        p[i] = (100-p[i]) / s[i]
    }

    var q Queue2 = Queue2{items: p}
    var days int = 0
    var cnt int

    for q.size() > 0{
        cnt = 0
        for {
            if q.size() == 0 {
                break
            }

            var i = q.peek()
            if i > days {
                break
            }

            cnt++
            q.pop()
        }

        if cnt != 0 {
            res = append(res, cnt)
        }

        days++
    }

    return res
}

func solution(p []int, s []int) (res []int) {
    var size int = len(p)
    var idx int = 0

    var cnt int = 0
    for i:=0; i<size; i++ {
        if p[i] < 100 {
            break
        }

        cnt++
        idx++
    }

    if cnt != 0 {
        res = append(res, cnt)
    }

    for {
        if idx == size {
            break
        }

        cnt = 0
        for i:=0; i<size; i++ {
            if i < idx {
                continue
            }

            if i == idx && p[i] > 99 {
                cnt++
                idx++
            }
            p[i] += s[i]
        }

        if cnt == 0 {
            continue
        }

        res = append(res, cnt)
    }

    return
}
