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

	var a, b, c, sum, mod int
	var arr []int = make([]int, 10)
	fmt.Fscanf(r, "%d\n", &a)
	fmt.Fscanf(r, "%d\n", &b)
	fmt.Fscanf(r, "%d\n", &c)

	sum = a * b * c

	for {
		mod = sum % 10
		arr[mod]++
		sum /= 10
		if sum == 0 {
			break
		}
	}

	for _, v := range arr {
		fmt.Fprintf(w, "%d\n", v)
	}
}
