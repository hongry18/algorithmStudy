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
	var n int64
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fprintln(w, fibo(n))
}

func fibo(x int64) int64 {
	if x < 2 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}
