package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()

	var N, x int

	fmt.Fscanf(r, "%d\n", &N)

	ans := 0
	for i := 0; i < N; i++ {
		fmt.Fscanf(r, "%d", &x)
		if sol(x) {
			ans += 1
		}
	}

	fmt.Fprintln(w, ans)
}

func sol(x int) bool {
	if x < 2 {
		return false
	}

	if x < 4 {
		return true
	}

	if x%2 == 0 || x%3 == 0 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
