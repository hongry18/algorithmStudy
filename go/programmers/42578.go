package main

import "fmt"

func main() {
    var a int
    //a=solution([][]string{{"yellow_hat", "headgear"}, {"blue", "headgear"}, {"gree_turban", "headgear"}})
    //a=solution([][]string{{"yellow_hat", "headgear"}, {"blue", "eyewear"}, {"gree_turban", "headgear"}})
    a=solution([][]string{{"yellow_hat", "1"}, {"blue", "2"}, {"gree_turban", "3"}})
    //a=solution([][]string{ {"yellow_hat", "headgear"}, {"blue", "a"}, {"gree_turban", "headgear"}, {"a", "a"}, {"bl", "a"}, {"a", "a"}, {"bl", "a"}, {"a", "c"}, {"bl", "c"}, {"a", "d"}, {"bl", "c"} })

    fmt.Println("answer: ", a)
}

func solution(c [][]string) int {
    var cats map[string]int = map[string]int{}
    var cnt int = 1
    var catCnt int = 0
    var items []int
    var length int

    for _, v := range c {
        _, exist := cats[v[1]]
        if !exist {
            cats[v[1]] = catCnt
            items = append(items, 0)
            catCnt++
        }

        items[cats[v[1]]]++
    }

    length = len(items)

    if length < 2 {
        return items[0]
    }

    for i:=0; i<length; i++ {
        cnt *= (items[i] + 1)
    }

    fmt.Println(cnt, items)

    return cnt - 1
}

func Comb(a []int, visit []bool, s, n, r int) int {
    if r == 0 {
        mul := 1
        for i:=0; i<n; i++ {
            if visit[i] {
                mul *= a[i]
            }
        }
        return mul
    }

    sum := 0
    for i:=s; i<n; i++ {
        visit[i] = true
        sum+=Comb(a, visit, i+1, n, r-1)
        visit[i] = false
    }

    return sum
}
