package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	const N = 202
	x := make([]int, N)
	y := make([]int, N)
	p := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i], &p[i])
	}

	var d [N][N]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dist := abs(x[i]-x[j]) + abs(y[i]-y[j])
			d[i][j] = (dist + p[i] - 1) / p[i]
		}
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				d[i][j] = min(d[i][j], max(d[i][k], d[k][j]))
			}
		}
	}

	ans := 4_000_000_000
	for i := 1; i <= n; i++ {
		t := 0
		for j := 1; j <= n; j++ {
			t = max(t, d[i][j])
		}
		ans = min(ans, t)
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
