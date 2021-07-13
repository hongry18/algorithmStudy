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

	var t int

	fmt.Fscanf(r, "%d\n", &t)

	for i := 0; i < t; i++ {
		var l int
		var s string
		fmt.Fscanf(r, "%d %s\n", &l, &s)

		for j := 0; j < len(s); j++ {
			for k := 0; k < l; k++ {
				fmt.Fprintf(w, "%s", string(s[j]))
			}
		}
		fmt.Fprintf(w, "\n")
	}
}
