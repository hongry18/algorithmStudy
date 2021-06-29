package main

import "fmt"

func main() {
    var a int
    a=solution("ABABAAAAABA")
    //a=solution("JAN")
    //a=solution("CANAAAAANAN")
    fmt.Println(a)
}

func solution(s string) int {
    var (
        items []byte = []byte(s)

        lc int = 0
        rc int = 0
        bc int = 0
        temp int
        id1 int
        id2 int = 100000
        cid int
        size int = len(items)
        back bool = false
    )

    for i, v := range items {
        temp = calc(v)
        if i < size / 2 && temp != 0 && int(calc(items[i+1])) == 0 {
            id1 = i
        }

        if i > size / 2 && temp != 0 && int(items[i-1]) == 0  {
            if i < id2 {
                id2 = i
            }
        }
        items[i] = byte(temp)
    }

    temp = 0
    for i,v := range items {
        if i==0 {
            rc += int(v)
            continue
        }
        if int(v) == 0 {
            temp++ 
            continue
        }
        rc += 1 + int(v) + temp
        temp = 0
    }

    cid = 0
    for {
        if id1 == 0 || id2 == 100000 {
            break
        }
        
        if cid == id2 {
            lc += int(items[cid])
            break
        }

        if cid == id1 {
            lc += int(items[cid])
            back = true
        }

        if !back && cid <= id1 {
            lc += int(items[cid])
        }

        if back && cid > id1 {
            lc += int(items[cid])
        }

        lc += 1

        if back {
            cid--
            if cid == -1 {
                cid = size - 1
            }
        } else {
            cid++
        }
    }

    bc = int(items[0])
    for i:=size-1; i>0; i-- {
        if int(items[i]) == 0 {
            temp++
            continue
        }
        bc += temp + int(items[i]) + 1
        temp = 0
    }

    //fmt.Println(id1, id2, lc, bc, rc)
    if id1 != 0 && id2 != 100000 && lc < rc {
        return lc
    }

    if bc < rc {
        return bc
    }
    return rc
}

func calc(c byte) int {
    r := int(c) - 65
    if r > 13 {
        r = 26 - r
    }
    return r
}
