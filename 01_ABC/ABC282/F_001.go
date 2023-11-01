package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)
	fmt.Fprintln(out, 12*n)
	out.Flush()
	for i := 0; i < 12; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fprintln(out, j, min(j+(1<<i)-1, n))
			out.Flush()
		}
	}
	var q int
	fmt.Scan(&q)
	for i := 1; i <= q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		x := log2(r - l + 1)
		fmt.Fprintln(out, n*x+l, n*x+r-(1<<x)+1)
		out.Flush()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}
