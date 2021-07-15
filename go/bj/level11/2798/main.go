package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	w = bufio.NewWriter(os.Stdout)
	r = bufio.NewReader(os.Stdin)
)

func main() {
	defer w.Flush()
	var N, M, C, size int
	var items []int

	fmt.Fscanf(r, "%d %d\n", &N, &M)

	for {
		_, err := fmt.Fscanf(r, "%d", &C)

		if err != nil {
			break
		}

		idx := sort.Search(size, func(i int) bool { return items[i] <= C })

		items = append(items, 0)
		copy(items[idx+1:], items[idx:])
		items[idx] = C
		size++
	}

	sum := items[N-3:][0] + items[N-3:][1] + items[N-3:][2]
	if N == 3 {
		fmt.Fprintln(w, sum)
		return
	}

	for i := 0; i < N-2; i++ {
		if items[i] >= M {
			continue
		}

		for j := i + 1; j < N-1; j++ {
			if items[i]+items[j] >= M {
				continue
			}

			for k := j + 1; k < N; k++ {
				s := items[i] + items[j] + items[k]
				if s > M {
					continue
				}
				if s > sum {
					sum = s
				}
			}
		}
	}

	fmt.Fprintln(w, sum)
}
