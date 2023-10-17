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
	var a [60]int
	var know, unkn [1010]int
	sum := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		know[i] = a[i]
		unkn[i] = n - 2
	}
	var m int
	fmt.Fscan(in, &m)
	var s [60][60]bool
	for m > 0 {
		m--
		var op, x, y int
		fmt.Fscan(in, &op, &x, &y)
		if op == 0 {
			s[x][y] = true
			know[x] += a[y]
			unkn[x]--
		} else {
			if s[x][y] {
				fmt.Fprintln(out, a[y], a[y])
			} else {
				fmt.Fprintln(out, max(0, sum-know[x]-unkn[x]*100), min(sum-know[x], 100))
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
