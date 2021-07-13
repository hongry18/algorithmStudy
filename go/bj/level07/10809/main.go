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

	var alphabets map[byte]int = make(map[byte]int)

	for i := 97; i < 123; i++ {
		alphabets[byte(i)] = -1
	}

	var s string
	fmt.Fscanf(r, "%s\n", &s)
	for i := 0; i < len(s); i++ {
		v, _ := alphabets[s[i]]
		if v != -1 {
			continue
		}

		alphabets[s[i]] = i
	}

	for i := 97; i < 123; i++ {
		fmt.Fprintf(w, "%d ", alphabets[byte(i)])
	}
}
