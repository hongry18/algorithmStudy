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

	var N int
	fmt.Fscanf(r, "%d", &N)

	for _, v := range Sol(N) {
		fmt.Fprintln(w, v)
	}
}

func Sol(x int) []int {
	if x < 2 {
		return nil
	}

	if x < 4 {
		return []int{x}
	}

	a := x
	var ans []int = make([]int, 0)
	for i := 2; i <= x; i++ {
		if x%i != 0 {
			continue
		}

		for {
			if a%i != 0 {
				break
			}

			a = a / i
			ans = append(ans, i)

			if a < 2 {
				break
			}

		}

		if a < 2 {
			break
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	return ans
}
