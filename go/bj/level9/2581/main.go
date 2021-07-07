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

	var M, N int
	fmt.Fscanf(r, "%d\n", &M)
	fmt.Fscanf(r, "%d\n", &N)

	var sum int = 0
	var min int = 10001
	for i := M; i <= N; i++ {
		if !IsPrime(i) {
			continue
		}

		sum += i
		if min > i {
			min = i
		}
	}

	if sum == 0 {
		fmt.Fprintln(w, -1)
		return
	}

	fmt.Fprintln(w, sum)
	fmt.Fprintln(w, min)
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

	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
