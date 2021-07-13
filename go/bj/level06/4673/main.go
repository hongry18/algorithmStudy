package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var cnt int = 10000

	var exists map[int]int = make(map[int]int)
	for i := 1; i <= cnt; i++ {
		exists[calc(i)] = 1
	}

	for i := 1; i <= cnt; i++ {
		if _, ok := exists[i]; ok {
			continue
		}

		fmt.Fprintf(w, "%d\n", i)
	}

}

func calc(i int) int {
	if i < 10 {
		return i + i
	}
	var a = i
	var sum int
	for {
		sum += i % 10
		i = i / 10
		if i == 0 {
			break
		}
	}
	return a + sum
}
