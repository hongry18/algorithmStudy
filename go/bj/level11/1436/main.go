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

	var N, x int = 2, 666

	if N == 1 {
		fmt.Fprintln(w, 666)
		return
	}

	for i := 1; i < N; i++ {

	}
}
