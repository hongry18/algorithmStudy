package main

import (
	"bufio"
	"bytes"
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

	arr := make([][]rune, N)
	for i := 0; i < N; i++ {
		arr[i] = make([]rune, N)
	}
	star(0, 0, N, false, arr)

	var b bytes.Buffer
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			b.WriteRune(arr[i][j])
		}
		b.WriteByte('\n')
	}

	fmt.Fprintln(w, b.String())
}

func star(x, y, N int, blank bool, arr [][]rune) {
	if blank {
		for i := x; i < x+N; i++ {
			for j := y; j < y+N; j++ {
				arr[i][j] = ' '
			}
		}
		return
	}

	if N == 1 {
		arr[x][y] = '*'
		return
	}

	size := N / 3
	cnt := 0

	for i := x; i < x+N; i += size {
		for j := y; j < y+N; j += size {
			cnt++

			star(i, j, size, cnt == 5, arr)
		}
	}
}
