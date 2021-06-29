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

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	var items []float64
	var max int
	var sum float64

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(r, "%d ", &x)

		if x > max {
			max = x
		}

		items = append(items, float64(x))
	}

	for i := 0; i < n; i++ {
		sum += items[i] / float64(max) * 100
	}

	fmt.Fprintf(w, "%f", sum/float64(n))
}
