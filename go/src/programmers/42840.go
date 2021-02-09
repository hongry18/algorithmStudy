package main

import (
    "os"
    "fmt"
    "sort"
    "bufio"
)

type item struct {
    key int
    val int
}

func main() {
    var answer1 = []int{1, 3, 2, 4, 2}
    //var answer1 = []int{1,2,3,4,5,1,2,3,4,5,3,1,2,3,5}
    solution(answer1)
}

func solution(answers []int) (list [] int) {

    var (
        a = []int{1, 2, 3, 4, 5}
        alen = len(a)
        b = []int{2, 1, 2, 3, 2, 4, 2, 5}
        blen = len(b)
        c = []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
        clen = len(c)
        ck = []item{{1,0}, {2,0}, {3,0}}
        _len = len(answers)
        _max = 0
        idx = 0
        fi = []item{}
        res = []int{}
    )
    
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    for i:=0; i<_len; i++ {
        idx = i % alen
        if answers[i] == a[idx] {
            ck[0].val++
        }

        idx = i % blen
        if answers[i] == b[idx] {
            ck[1].val++
        }

        idx = i % clen
        if answers[i] == c[idx] {
            ck[2].val++
        }
    }

    for i:=0; i<3; i++ {
        if ck[i].val > _max {
            _max = ck[i].val
        }
    }

    for i:=0; i<3; i++ {
        if ck[i].val != _max {
            continue
        }
        fi = append(fi, ck[i])
    }

    f := func(x, y int) bool {
        return fi[x].key < fi[y].key
    }

    sort.Slice(fi, f)

    for i:=0; i<len(fi); i++ {
        res = append(res, fi[i].key)
    }

    fmt.Fprintf(writer, "%v %v %d", fi, res, _max)

    return
}
