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

	var s string
	var alphabets []int = make([]int, 26)
	var max int
	var equals bool
	var out int
	fmt.Fscanf(reader, "%s\n", &s)

	for i := 0; i < len(s); i++ {
		b := s[i]
		if b < 96 {
			b += 32
		}

		alphabets[b-97]++
	}

	for i := 0; i < 26; i++ {
		if alphabets[i] > max {
			max = alphabets[i]
			out = i
			equals = false
		} else if alphabets[i] == max && max != 0 {
			equals = true
		}
	}

	if equals {
		fmt.Fprintf(writer, "?")
	} else {
		fmt.Fprintf(writer, "%s", string(byte(out)+65))
	}
}
