package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	defer writer.Flush()
	var x, y, w, h int
	var ans []int

	fmt.Fscanf(reader, "%d %d %d %d", &x, &y, &w, &h)
	ans = append(ans, x)
	ans = append(ans, y)
	ans = append(ans, w-x)
	ans = append(ans, h-y)
	sort.Ints(ans)

	fmt.Fprintln(writer, ans[0])
}
