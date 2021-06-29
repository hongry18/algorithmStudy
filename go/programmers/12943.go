package main

func main() {
    solution(6)
    solution(16)
}

func solution(n int) int {
    if n == 1 {
        return 0
    }
    var cnt = 0
    for {
        if cnt == 500 {
            cnt = -1
            break
        }
        if n % 2 == 0 {
            n /= 2
        } else {
            n = n * 3 + 1
        }

        cnt++
        if n == 1 {
            break
        }
    }
    return cnt
}
