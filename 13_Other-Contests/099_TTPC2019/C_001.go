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

	var n, x int
	fmt.Fscan(in, &n, &x)

	a := make([]int, n)
	y := x
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] >= 0 {
			y ^= a[i]
		}
	}
	for i := 0; i < n; i++ {
		if a[i] == -1 {
			a[i] = min(x, y)
			y ^= a[i]
		}
	}
	if y == 0 {
		for _, i := range a {
			fmt.Fprintf(out, "%d ", i)
		}
	} else {
		fmt.Fprintln(out, -1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
