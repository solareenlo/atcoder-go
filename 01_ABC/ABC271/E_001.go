package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	var x, y, w [N]int
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &x[i], &y[i], &w[i])
	}

	var f [N]int
	for i := range f {
		f[i] = 1_000_000_000_000_000_000
	}
	f[1] = 0
	for i := 1; i <= k; i++ {
		var s int
		fmt.Fscan(in, &s)
		f[y[s]] = min(f[x[s]]+w[s], f[y[s]])
	}

	if f[n] == f[0] {
		fmt.Println(-1)
	} else {
		fmt.Println(f[n])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
