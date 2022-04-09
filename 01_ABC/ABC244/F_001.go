package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	const N = 20
	g := [N][N]int{}
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		g[x][y] = 1
		g[y][x] = 1
	}

	f := [1 << 17][N]int{}
	type pair struct{ x, y int }
	q := make([]pair, 0)
	for i := 1; i <= n; i++ {
		f[0][i] = 1
		q = append(q, pair{0, i})
	}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for i := 1; i <= n; i++ {
			if g[p.y][i] != 0 && f[p.x^(1<<i-1)][i] == 0 {
				f[p.x^(1<<i-1)][i] = f[p.x][p.y] + 1
				q = append(q, pair{p.x ^ (1<<i - 1), i})
			}
		}
	}

	ans := 0
	for i := 0; i < (1 << n); i++ {
		res := 1 << 60
		for j := 1; j <= n; j++ {
			res = min(res, f[i][j]-1)
		}
		ans += res
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
