package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	var n, x, max, min int
	max = -1000000
	min = 1000000

	fmt.Fscanf(r, "%d\n", &n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanf(r, "%d", &x)
		} else {
			fmt.Fscanf(r, "%d ", &x)
		}

		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Fprintf(w, "%d %d\n", min, max)
}
