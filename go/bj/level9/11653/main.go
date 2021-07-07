package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()

	// var N int
	// fmt.Fscanf(r, "%d", &N)
	Prime(72)
}

func Prime(x int) []int {
	if x < 4 {
		return []int{x}
	}

	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			fmt.Println(x, i, x%i)
		}
	}

	return []int{}
}
