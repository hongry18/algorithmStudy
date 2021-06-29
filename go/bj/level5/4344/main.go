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

	var c int
	fmt.Fscanf(r, "%d\n", &c)

	for i := 0; i < c; i++ {
		var n int
		var sum int
		var cnts []int
		var avg int
		var gt_cnt int
		fmt.Fscanf(r, "%d", &n)

		for j := 0; j < n; j++ {
			var userCnt int
			fmt.Fscanf(r, " %d", &userCnt)

			sum += userCnt
			cnts = append(cnts, userCnt)
		}
		fmt.Fscanf(r, "\n")

		avg = sum / n

		for _, v := range cnts {
			if v <= avg {
				continue
			}

			gt_cnt += 1
		}

		res := (float64(gt_cnt) / float64(n)) * 100
		fmt.Fprintf(w, "%.3f%%\n", res)
	}
}
