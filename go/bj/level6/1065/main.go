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

	var x int
	fmt.Fscanf(r, "%d\n", &x)
	fmt.Fprintf(w, "%d\n", cal(x))
}

func cal(x int) int {
	if x < 100 {
		return x
	}

	var res int = 99

	for i := 100; i <= x; i++ {
		if !check(i) {
			continue
		}

		res += 1
	}

	return res
}

func check(x int) bool {
	f1 := x % 10
	x /= 10
	f2 := x % 10
	f := f2 - f1

	if f > 4 {
		return false
	}

	x /= 10

	for {
		c := x % 10
		if (f2 + f) != c {
			return false
		}

		x /= 10
		f2 = c
		if x == 0 {
			break
		}
	}
	return true
}
