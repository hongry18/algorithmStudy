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
	var M, N int
	fmt.Fscanf(r, "%d %d\n", &M, &N)

	for i := M; i <= N; i++ {
		if !IsPrime(i) {
			continue
		}

		fmt.Fprintln(w, i)
	}
}

func IsPrime(x int) bool {
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
