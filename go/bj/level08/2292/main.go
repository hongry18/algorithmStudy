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
	var a int = 0
	fmt.Fscanf(r, "%d\n", &n)

	for {
		if 1+((a*(a*6+6))/2) >= n {
			break
		}

		a++
	}

	fmt.Fprintln(w, a+1)

}
