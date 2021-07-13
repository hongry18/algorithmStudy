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

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	fmt.Fprintln(w, calc(n))
}

func calc(n int) int {
	var ans int

	mod := n % 5

	if mod == 0 {
		ans = n / 5
	} else if mod == 1 {
		// 6, 11, 16, 21
		ans = (n / 5) + 1
	} else if mod == 2 && n >= 12 {
		// 12 17 22 27
		ans = (n / 5) + 2
	} else if mod == 3 {
		ans = (n / 5) + 1
	} else if mod == 4 && n >= 9 {
		// 9 14 19
		ans = (n / 5) + 2
	} else {
		ans = -1
	}

	return ans
}
