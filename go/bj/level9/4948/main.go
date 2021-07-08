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

	var size int = 123470 * 2
	var prime []int = make([]int, size)
	prime[0] = 1
	prime[1] = 1

	for i := 2; i*i <= size; i++ {
		if prime[i] != 0 {
			continue
		}

		for j := i * i; j < size; j += i {
			prime[j] = 1
		}
	}

	for {
		var n int
		fmt.Fscanf(r, "%d\n", &n)
		if n == 0 {
			break
		}

		var cnt int
		for i := n + 1; i <= n*2; i++ {
			if prime[i] == 0 {
				cnt++
			}
		}

		fmt.Fprintln(w, cnt)
	}
}
