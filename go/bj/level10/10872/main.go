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

	var N int

	fmt.Fscanf(r, "%d\n", &N)

	fmt.Fprintln(w, Factorial(N))
}

func Factorial(x int) int64 {
	var ans int64 = 1

	for i := 1; i <= x; i++ {
		ans *= int64(i)
	}

	return ans
}
