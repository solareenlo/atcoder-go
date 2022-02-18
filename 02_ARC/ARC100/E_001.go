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

	m := 1 << n
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i])
	}

	r := make([]int, m)
	g := make([]int, m)
	for i := 0; i < m; i++ {
		for j := i; j < m; j = (j + 1) | i {
			r[j] = max(r[j], g[j]+a[i])
			g[j] = max(g[j], a[i])
		}
	}

	ans := 0
	for i := 1; i < m; i++ {
		ans = max(ans, r[i])
		fmt.Fprintln(out, ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
