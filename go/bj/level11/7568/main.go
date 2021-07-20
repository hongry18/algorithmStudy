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

	var N, x, y int
	var ans [][]int = make([][]int, 0)
	fmt.Fscanf(r, "%d\n", &N)

	for i := 0; i < N; i++ {
		fmt.Fscanf(r, "%d %d\n", &x, &y)

		ans = append(ans, []int{x, y, 0, i, 0})
	}

	fmt.Println()
	fmt.Println()
	for i := 0; i < N; i++ {
		var cnt int = 1
		for j := 0; j < N; j++ {
			fmt.Println(i, j)
			if ans[i][0] > ans[j][0] && ans[i][1] > ans[j][1] {
				cnt += 1
			}
		}

		fmt.Fprintf(w, "%d ", cnt)
	}
}
