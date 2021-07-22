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
	var N, M, min int

	min = 100000000

	fmt.Fscanf(r, "%d %d\n", &N, &M)
	var maps [][]int = make([][]int, N)

	var whiteMap [][]int = make([][]int, 8)
	var blackMap [][]int = make([][]int, 8)

	for i := 0; i < 8; i++ {
		whiteMap[i] = make([]int, 8)
		blackMap[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			if j == 0 {
				whiteMap[i][j] = i % 2
			} else {
				if whiteMap[i][j-1] == 0 {
					whiteMap[i][j] = 1
				} else {
					whiteMap[i][j] = 0
				}
			}

			if whiteMap[i][j] == 0 {
				blackMap[i][j] = 1
			} else {
				blackMap[i][j] = 0
			}
		}
	}

	for i := 0; i < N; i++ {
		maps[i] = make([]int, M)
		var line string
		fmt.Fscanf(r, "%s\n", &line)
		for j := 0; j < M; j++ {
			if line[j] == 'B' {
				maps[i][j] = 0
			} else {
				maps[i][j] = 1
			}
		}
	}

	for i := 0; i <= N-8; i++ {
		for j := 0; j <= M-8; j++ {
			_m := cmp(maps, i, j, whiteMap)

			if _m < min {
				min = _m
			}

			if min == 0 {
				fmt.Fprintln(w, min)
				return
			}

			_m = cmp(maps, i, j, blackMap)

			if _m < min {
				min = _m
			}

			if min == 0 {
				fmt.Fprintln(w, min)
				return
			}
		}
	}

	fmt.Fprintln(w, min)
}

func cmp(a [][]int, x, y int, b [][]int) int {
	var ans int = 0

	for i := x; i < x+8; i++ {
		for j := y; j < y+8; j++ {
			if a[i][j] == b[i%8][j%8] {
				continue
			}

			ans += 1
		}
	}

	return ans
}
