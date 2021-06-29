package main

import "fmt"

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 6}))
}

func sum(a []int) int64 {
	var r int64
	for i := 0; i < len(a); i++ {
		r += int64(a[i])
	}
	return r
}
