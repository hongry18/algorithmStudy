package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()

	var n int

	fmt.Fscanf(r, "%d\n", &n)

	var s string
	var sum int64

	fmt.Fscanf(r, "%s\n", &s)
	for i := 0; i < n; i++ {
		x, e := strconv.ParseInt(string(s[i]), 10, 64)
		if e != nil {
			continue
		}
		sum += x
	}

	fmt.Fprintf(w, "%d\n", sum)
}
