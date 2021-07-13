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

	var a, b, c int

	for {
		fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		if IsTriangle(a, b, c) {
			fmt.Fprintln(w, "right")
		} else {
			fmt.Fprintln(w, "wrong")
		}
	}
}

func IsTriangle(a, b, c int) bool {
	a *= a
	b *= b
	c *= c

	if a+b == c || a+c == b || b+c == a {
		return true
	}

	return false
}
