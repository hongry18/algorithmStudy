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
	var N int
	fmt.Fscanf(r, "%d\n", &N)
	var ans [][]int = make([][]int, 0)
	hanoi(N, 1, 3, 2, &ans)

	size := len(ans)

	fmt.Fprintln(w, size)

	for i := 0; i < size; i++ {
		fmt.Fprintf(w, "%d %d\n", ans[i][0], ans[i][1])
	}
}

func hanoi(x, from, to, tmp int, ans *[][]int) {
	if x == 0 {
		return
	}

	hanoi(x-1, from, tmp, to, ans)
	*ans = append(*ans, []int{from, to})
	hanoi(x-1, tmp, to, from, ans)
}
