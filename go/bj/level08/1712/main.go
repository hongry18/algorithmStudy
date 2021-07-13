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

/*
손익분기점 계산법
고정영업비용 / (판매가 - 단위당 변동비)
*/

func main() {
	defer writer.Flush()

	var a, b, c int64
	fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)

	if b >= c {
		fmt.Fprintln(writer, "-1")
		return
	}

	fmt.Fprintln(writer, a/(c-b)+1)
}
