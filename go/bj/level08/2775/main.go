package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()

	// var T, k, n int
	var m [][]int = make([][]int, 15)

	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			if i == 0 {
				m[i] = append(m[i], j+1)
				continue
			}

			if j == 0 {
				m[i] = append(m[i], j+1)
				continue
			}

			m[i] = append(m[i], m[i][j-1]+m[i-1][j])
		}
	}

	j, _ := json.MarshalIndent(m, "", "  ")
	fmt.Fprintln(w, string(j))
	// fmt.Fscanf(r, "%d\n", &T)

	// for i := 0; i < T; i++ {
	// 	fmt.Fscanf(r, "%d\n", &k)
	// 	fmt.Fscanf(r, "%d\n", &n)

	// 	fmt.Fprintln(w, m[k][n-1])
	// }
}
