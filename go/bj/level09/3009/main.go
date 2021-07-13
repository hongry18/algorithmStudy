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

	var maps [][]int = make([][]int, 0)

	for i := 0; i < 3; i++ {
		var p1, p2 int
		_, e := fmt.Fscanf(r, "%d %d\n", &p1, &p2)
		if e != nil {
			break
		}
		maps = append(maps, []int{p1, p2})
	}

	var x, y int

	if maps[0][0] == maps[2][0] {
		x = maps[1][0]
	} else if maps[1][0] == maps[2][0] {
		x = maps[0][0]
	} else {
		x = maps[2][0]
	}

	if maps[0][1] == maps[2][1] {
		y = maps[1][1]
	} else if maps[1][1] == maps[2][1] {
		y = maps[0][1]
	} else {
		y = maps[2][1]
	}

	fmt.Fprintln(w, x, y)
}
