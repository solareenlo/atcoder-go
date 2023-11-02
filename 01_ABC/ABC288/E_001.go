package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 5555

	var a, c [N]int
	var vis [N]bool
	var f [N][N]int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = int(2e18)
		}
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}
	for i := 1; i <= m; i++ {
		var x int
		fmt.Fscan(in, &x)
		vis[x] = true
	}
	f[0][0] = 0
	for i := 1; i <= n; i++ {
		Min := c[i]
		for j := 0; j < i+1; j++ {
			Min = min(Min, c[i-j])
			if !vis[i] {
				f[i][j] = min(f[i][j], f[i-1][j])
			}
			if j < i {
				f[i][j+1] = f[i-1][j] + a[i] + Min
			}
		}
	}
	ans := int(2e18)
	for i := m; i <= n; i++ {
		ans = min(ans, f[n][i])
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
