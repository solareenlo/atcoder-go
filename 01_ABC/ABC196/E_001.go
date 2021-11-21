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

	var n int
	fmt.Fscan(in, &n)

	A, B, C := -1<<60, 1<<60, 0
	for i := 1; i < n+1; i++ {
		var a, t int
		fmt.Fscan(in, &a, &t)
		if t == 1 {
			A += a
			B += a
			C += a
		} else if t == 2 {
			A = max(A, a)
			B = max(B, a)
		} else {
			A = min(A, a)
			B = min(B, a)
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		fmt.Fprintln(out, min(B, max(A, x+C)))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
