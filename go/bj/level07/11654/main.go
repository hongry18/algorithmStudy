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

	var n string

	fmt.Fscanf(r, "%s\n", &n)

	fmt.Fprintf(w, "%d", n[0])
}
