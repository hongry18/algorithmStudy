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
	var ans int = 0
	fmt.Fscanf(r, "%d\n", &N)

	if N < 11 {
		fmt.Fprintln(w, 0)
		return
	}

	length := size(N)
	loop_range := length * 10

	for i := N - loop_range; i <= N; i++ {
		if i+ds(i) != N {
			continue
		}
		ans = i
		break
	}

	fmt.Fprintln(w, ans)
}

func ds(x int) int {
	s := 0

	for {
		s += x % 10
		x /= 10

		if x == 0 {
			break
		}
	}

	return s
}

func size(x int) int {
	s := 1

	if x < 10 {
		return s
	}

	for {
		x /= 10
		s++

		if x < 10 {
			return s
		}
	}
}
