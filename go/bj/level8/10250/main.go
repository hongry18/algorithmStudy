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

	// T 문제수, H 높이, W 넓이, N 번째 손님
	var T, H, W, N int

	fmt.Fscanf(r, "%d\n", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(r, "%d %d %d\n", &H, &W, &N)
		calc(w, H, W, N)
	}

}

func calc(w *bufio.Writer, H int, W int, N int) {
	X := int64(math.Ceil(float64(N) / float64(H)))
	Y := N % H

	if Y == 0 {
		Y = H
	}

	if X < 10 {
		fmt.Fprintf(w, "%d0%d\n", Y, X)
	} else {
		fmt.Fprintf(w, "%d%d\n", Y, X)
	}
}
