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

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 10000005
	a := make([]int, N)
	b := make([]int, N)
	p := make([]int, N)
	l, r := m+1, 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		r = max(r, a[i])
		l = min(l, b[i])
		p[a[i]] = max(p[a[i]], b[i])
	}

	ans := make([]int, N)
	for i := 1; i <= l; i++ {
		ans[r-i+1]++
		ans[m-i+2]--
		r = max(r, p[i])
	}

	for i := 1; i <= m; i++ {
		ans[i] = ans[i] + ans[i-1]
		fmt.Fprintf(out, "%d ", ans[i])
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
