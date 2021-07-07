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

	var a, b string
	fmt.Fscanf(r, "%s %s\n", &a, &b)

	fmt.Fprintln(w, sol(a, b))
}

func sol(a, b string) string {
	ba := reverse([]byte(a))
	bb := reverse([]byte(b))
	a_len := len(ba)
	b_len := len(bb)
	var res []byte = make([]byte, 0)
	var over bool = false

	if a_len < b_len {
		ba, bb = bb, ba
		a_len, b_len = b_len, a_len
	}

	for i := 0; i < b_len; i++ {
		s := bb[i] + ba[i] - 96
		if over {
			s += 1
		}

		if s < 10 {
			res = append(res, s+48)
			over = false
		} else {
			res = append(res, s+38)
			over = true
		}

	}

	if a_len == b_len {
		if over {
			res = append(res, 49)
		}
	} else {
		for i := b_len; i < a_len; i++ {
			if over {
				if ba[i]+1 == 58 {
					res = append(res, 48)
					over = true
				} else {
					res = append(res, ba[i]+1)
					over = false
				}

				continue
			}

			res = append(res, ba[i])
		}

		if over {
			res = append(res, 49)
		}
	}

	return string(reverse(res))
}

func reverse(data []byte) []byte {
	size := len(data)
	for i := 0; i < size/2; i++ {
		j := size - i - 1
		data[i], data[j] = data[j], data[i]
	}
	return data
}
