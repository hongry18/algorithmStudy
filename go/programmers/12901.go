package main

import (
    "./bp"
)

func main() {
    solution(5, 24)
}

func solution(a, b int) string {
    days := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
    months := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

    sum := 4 + b

    for i:=0; i<a-1; i++ {
        sum += months[i]
    }

    return days[sum%7]
}
