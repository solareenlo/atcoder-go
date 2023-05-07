package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	var X, Y [1 << 17]int
	var lx, rx, ly, ry int
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		X[i] = x + y
		Y[i] = x - y
		if i == 0 {
			lx = X[i]
			rx = X[i]
			ly = Y[i]
			ry = Y[i]
		}
		lx = min(lx, X[i])
		rx = max(rx, X[i])
		ly = min(ly, Y[i])
		ry = max(ry, Y[i])
	}
	for Q > 0 {
		Q--
		var id int
		fmt.Fscan(in, &id)
		id--
		fmt.Fprintln(out, max(rx-X[id], X[id]-lx, ry-Y[id], Y[id]-ly))
	}
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
