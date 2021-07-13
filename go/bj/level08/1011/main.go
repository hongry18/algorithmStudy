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

	var T, x, y int
	fmt.Fscanf(r, "%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(r, "%d %d\n", &x, &y)
		fmt.Fprintln(w, sol(x, y))
	}
}

func sol(x, y int) uint64 {
	s := y - x
	if s < 3 {
		return uint64(s)
	}

	i := 2
	d := 2
	var cnt uint64 = 2
	for {
		i += d
		cnt += 1

		if cnt%2 == 0 {
			d += 1
		}

		if i >= s {
			break
		}
	}

	return cnt
}
