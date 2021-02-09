package main

func main() {
    solution(3)
    solution(4)
}

func solution(n int) string {
    if n % 2 == 0 {
        return "Even"
    }

    return "Odd"
}
