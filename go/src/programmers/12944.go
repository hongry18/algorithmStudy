package main

func main() {
    solution([]int{1,2,3,4})
}

func solution(a []int) float64 {
    var size int = len(a)
    if size == 1 {
        return float64(a[0])
    }

    var sum float64 = 0
    for _, v := range a {
        sum += float64(v)
    }

    return sum / float64(size)
}
