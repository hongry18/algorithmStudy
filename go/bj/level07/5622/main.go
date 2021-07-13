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

	var sum int
	for _, v := range s {
		switch v {
		case 65, 66, 67:
			sum += 3
		case 68, 69, 70:
			sum += 4
		case 71, 72, 73:
			sum += 5
		case 74, 75, 76:
			sum += 6
		case 77, 78, 79:
			sum += 7
		case 80, 81, 82, 83:
			sum += 8
		case 84, 85, 86:
			sum += 9
		case 87, 88, 89, 90:
			sum += 10
		}
	}

	fmt.Fprintln(writer, sum)
}
