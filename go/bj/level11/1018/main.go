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
	var N, M int

	fmt.Fscanf(r, "%d %d\n", &N, &M)
	var maps [][]int = make([][]int, N)

	var whiteMap [][]int = make([][]int, 8)
	var blackMap [][]int = make([][]int, 8)

	for i := 0; i < 8; i++ {
		whiteMap[i] = make([]int, 8)
		blackMap[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			if j%2 == 0 {
				whiteMap[i][j] = 1
				blackMap[i][j] = 0
			} else {
				whiteMap[i][j] = 0
				blackMap[i][j] = 1
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

	// fmt.Println(maps)
}
