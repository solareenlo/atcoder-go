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

	const N = 3030

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var f [N][N][2]int
	for i := range f {
		for j := range f[i] {
			for k := range f[i][j] {
				f[i][j][k] = int(1e18)
			}
		}
	}
	f[0][0][1] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j][0] = min(f[i-1][j][0], f[i-1][j][1]+1)
			if j >= a[i] {
				f[i][j][1] = min(f[i-1][j-a[i]][0], f[i-1][j-a[i]][1])
			}
		}
	}

	for i := 1; i <= m; i++ {
		t := min(f[n][i][0], f[n][i][1])
		if t > n {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, t)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
