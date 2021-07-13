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

	var a, b, c int

	for {
		b = 1 + ((a * (a + 3)) / 2)
		if b >= x {
			break
		}

		a++
	}

	c = b - x

	d1 := 1 + c
	d2 := a + 1 - c

	if a%2 == 0 {
		fmt.Fprintf(w, "%d/%d", d1, d2)
	} else {
		fmt.Fprintf(w, "%d/%d", d2, d1)
	}
}
