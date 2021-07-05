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
	var a, b, v int

	fmt.Fscanf(r, "%d %d %d\n", &a, &b, &v)

	if a >= v {
		fmt.Fprintln(w, 1)
		return
	}

	fmt.Fprintln(w, (v-b-1)/(a-b)+1)
}
