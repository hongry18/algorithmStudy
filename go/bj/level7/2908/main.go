package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	defer writer.Flush()

	var a, b, ra, rb string

	fmt.Fscanf(reader, "%s %s\n", &a, &b)
	ra = reverse(a)
	rb = reverse(b)

	if ra > rb {
		fmt.Fprintf(writer, "%s\n", ra)
	} else {
		fmt.Fprintf(writer, "%s\n", rb)
	}

}

func reverse(s string) string {
	var result string
	for _, v := range s {
		result = strings.Join([]string{string(v), result}, "")
	}

	return result
}
