package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()

	var items map[int]int = make(map[int]int)

	for i := 0; i < 10; i++ {
		var n int
		fmt.Fscanf(r, "%d\n", &n)

		d := n % 42

		if _, ok := items[d]; ok {
			items[d]++
		} else {
			items[d] = 1
		}
	}

	fmt.Fprintf(w, "%d\n", len(items))
}
