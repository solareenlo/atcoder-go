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

	const N = 262149
	const M = 23

	var n, v int
	fmt.Fscan(in, &n, &v)

	j := v
	x := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}

	u := make([]int, N)
	m := 0
	for v != 0 {
		v /= 2
		u[m] = v
		m++
	}

	u[m] = j
	o := 1 << m
	l := [M][N]int{}
	r := [M][N]int{}
	for i := 0; i <= m; i++ {
		l[i][1] = 1
		r[i][n] = n
		for j := 2; j <= n; j++ {
			if x[j]-x[j-1] > u[i] {
				l[i][j] = j
			} else {
				l[i][j] = l[i][j-1]
			}
		}
		for j := n - 1; j > 0; j-- {
			if x[j+1]-x[j] > u[i] {
				r[i][j] = j
			} else {
				r[i][j] = r[i][j+1]
			}
		}
	}

	g := make([]int, N)
	for i := 0; i < o; i++ {
		g[i] = n + 1
	}
	f := make([]int, N)
	for i := 0; i < o; i++ {
		for j := 0; j < m; j++ {
			if i>>j&1 != 0 {
				f[i] = max(f[i], r[j][min(n, f[i^(1<<j)]+1)])
				g[i] = min(g[i], l[j][max(1, g[i^(1<<j)]-1)])
			}
		}
	}

	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = -1
	}
	for i := 0; i < o; i++ {
		a[g[i]] = max(a[g[i]], f[o-1^i])
	}

	for i := 1; i <= n; i++ {
		a[i+1] = max(a[i+1], a[i])
	}

	for i := 1; i <= n; i++ {
		if a[r[m][i]+1] >= l[m][i]-1 {
			fmt.Fprintln(out, "Possible")
		} else {
			fmt.Fprintln(out, "Impossible")
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
