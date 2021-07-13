package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var R float64

	var U, D float64 = math.Pi, 2

	fmt.Fscanf(r, "%f\n", &R)

	R *= R

	fmt.Fprintf(w, "%.6f\n", R*U)
	fmt.Fprintf(w, "%.6f\n", R*D)
}
