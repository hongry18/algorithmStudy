package main

import "fmt"

func main() {
    //solution(2, 10, []int{7,4,5,6})
    //solution(100, 100, []int{10})
    solution(100, 100, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10})
}

type Queue struct {
    i []int
}

func (q *Queue) push(d int) {
    q.i = append(q.i, d)
}

func (q *Queue) pop() (res int) {
    if len(q.i) < 1 {
        return 0
    }
    res = q.i[0]
    q.i = q.i[1:]
    return
}

func (q *Queue) size() int {
    return len(q.i)
}

func (q *Queue) peek() int {
    if len(q.i) < 1 {
        return 0
    }
    return q.i[0]
}

func solution(b_len, weight int, t_weights []int) int {

    // 다리
    b := Queue{i: make([]int, b_len)}
    // 대기중
    l := Queue{i: t_weights}
    // 골
    g := Queue{}

    var size int = l.size()

    var times int = 0
    var tmp int = 0
    var w int = 0
    var end int = 0

    for g.size() != size {

        tmp = l.peek()

        end = b.pop()
        if end != 0 {
            g.push(end)
        }
        w -= end

        if w + tmp <= weight {
            l.pop()
        } else {
            tmp = 0
        }
        w += tmp

        b.push(tmp)
        fmt.Println(b)
        times++
    }

    fmt.Println(times)

    return times
}
