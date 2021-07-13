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
	var T, x1, y1, r1, x2, y2, r2 int

	fmt.Fscanf(r, "%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(r, "%d %d %d %d %d %d\n", &x1, &y1, &r1, &x2, &y2, &r2)

		d := Pow2(x2-x1) + Pow2(y2-y1)

		if d == 0 && r1 == r2 {
			fmt.Fprintln(w, -1)
			continue
		}

		if Pow2(r2-r1) < d && Pow2(r2+r1) > d {
			fmt.Fprintln(w, 2)
			continue
		}

		if Pow2(r2-r1) == d || Pow2(r2+r1) == d {
			fmt.Fprintln(w, 1)
			continue
		}

		fmt.Fprintln(w, 0)
	}
}

func Pow2(x int) int {
	return x * x
}
