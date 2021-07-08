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

	prime := Prime()
	var T, n int
	fmt.Fscanf(r, "%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(r, "%d\n", &n)
		if n%2 == 1 {
			continue
		}

		d := n / 2

		if prime[d] == 0 {
			fmt.Fprintf(w, "%d %d\n", d, d)
			continue
		}

		for {
			d -= 1

			if prime[d] == 1 {
				continue
			}

			if prime[n-d] == 1 {
				continue
			}

			fmt.Fprintf(w, "%d %d\n", d, n-d)
			break
		}
	}
}

func Prime() []int {
	var x int = 10000
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
