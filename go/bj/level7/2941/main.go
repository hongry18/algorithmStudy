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
	fmt.Fscanf(reader, "%s\n", &s)

	size := len(s) - 1

	if size < 1 {
		fmt.Fprintf(writer, "%d\n", size+1)
		return
	}

	// = is 61
	// - is 45
	// j is 106

	var cnt int
	var idx int

	for {
		switch s[idx] {
		case 99:
			// c
			if idx+1 <= size && (s[idx+1] == 61 || s[idx+1] == 45) {
				idx += 2
			} else {
				idx += 1
			}
			cnt += 1
		case 100:
			// d
			if idx+1 <= size && s[idx+1] == 45 {
				idx += 2
			} else if idx+2 <= size && s[idx+1] == 122 && s[idx+2] == 61 {
				idx += 3
			} else {
				idx += 1
			}
			cnt += 1
		case 108, 110:
			// l
			if idx+1 <= size && s[idx+1] == 106 {
				idx += 2
			} else {
				idx += 1
			}

			cnt += 1
		case 115, 122:
			// s, z
			if idx+1 <= size && s[idx+1] == 61 {
				idx += 2
			} else {
				idx += 1
			}

			cnt += 1
		default:
			// other
			idx += 1
			cnt += 1
		}

		if idx > size {
			break
		}
	}

	fmt.Fprintf(writer, "%d\n", cnt)
}
