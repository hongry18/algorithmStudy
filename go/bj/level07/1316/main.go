package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	var cnt int

	for i := 0; i < n; i++ {
		var alphabets string
		fmt.Fscanf(reader, "%s\n", &alphabets)
		if isGroup(alphabets) {
			cnt += 1
		}
	}

	fmt.Fprintf(writer, "%d\n", cnt)
}

func isGroup(s string) bool {
	var dup []int = make([]int, 27)

	var prev rune
	for _, v := range s {
		cur := v - 96

		if prev != cur {
			if dup[cur] == 1 {
				return false
			}
			dup[cur] = 1
		}

		prev = cur
	}

	return true
}
