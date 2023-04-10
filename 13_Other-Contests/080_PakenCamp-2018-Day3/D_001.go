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
	var c, Pre, Suc [100010]int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}
	for i := 2; i <= n; i++ {
		Pre[i] = max(Pre[i-1]+c[i-1], 0)
	}
	for i := n - 1; i >= 1; i-- {
		Suc[i] = max(Suc[i+1]+c[i], 0)
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, max(Pre[i], Suc[i]))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
