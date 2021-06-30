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

	s, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	var prev rune
	var cnt int = 0

	for _, v := range s {
		if v > 96 {
			v -= 32
		}

		if (prev == 0 || prev == 32) && v > 64 && v < 91 {
			cnt++
		}

		prev = v
	}

	fmt.Fprintf(writer, "%d\n", cnt)
}
