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
	fmt.Fscanf(r, "%d %d\n", &M, &N)

	prime := Prime(N)

	for i := M; i <= N; i++ {
		if prime[i] == 1 {
			continue
		}

		fmt.Fprintln(w, i)
	}
}

func Prime(x int) []int {
	var prime []int = make([]int, x+1)
	prime[0] = 1
	prime[1] = 1

	for i := 2; i*i <= x; i++ {
		if prime[i] != 0 {
			continue
		}

		for j := i * i; j <= x; j += i {
			prime[j] = 1
		}
	}

	return prime
}
