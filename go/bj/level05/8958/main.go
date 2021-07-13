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

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanf(r, "%s\n", &s)

		var sum int
		var val int
		for j := 0; j < len(s); j++ {
			if s[j] == 79 {
				val += 1
			} else {
				val = 0
			}

			sum += val
		}

		fmt.Fprintf(w, "%d\n", sum)
	}
}
