package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 300007
const MOD = 998244353
const I2 = MOD - MOD/2

var (
	m    int
	len  int
	root int
	e    = make([][]int, N)
	fa   = make([]int, N)
	f    = [N][3][3]int{}
)

func inc(a, b int) int {
	a += b - MOD
	a += a >> 31 & MOD
	return a
}

func mul(a, b int) int { return a * b % MOD }

func DFS(u, dep int) {
	if dep > len {
		len = dep
		root = u
	}
	for _, v := range e[u] {
		if v^fa[u] != 0 {
			fa[v] = u
			DFS(v, dep+1)
		}
	}
}

func dp(u, fa, d int) {
	if d == m/2 {
		f[u][1][1] = 1
	} else {
		f[u][0][0] = 1
	}
	for _, v := range e[u] {
		if v == fa {
			continue
		}
		dp(v, u, d+1)
		g := [3][3]int{}
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				for p := 0; p < 3; p++ {
					for q := 0; q < 3; q++ {
						for k := 0; k < 3; k++ {
							tmp := 0
							if k == 1 {
								tmp = 1
							}
							r := min(2, x+p*tmp)
							tmp = 0
							if k == 2 {
								tmp = 1
							}
							t := min(2, y+q*tmp)
							g[r][t] = inc(g[r][t], mul(f[u][x][y], f[v][p][q]))
						}
					}
				}
			}
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				f[u][i][j] = g[i][j]
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}

	DFS(1, 1)

	for i := range fa {
		fa[i] = 0
	}
	len = 0
	DFS(root, 1)

	dia := make([]int, N)
	for i := root; i > 0; i = fa[i] {
		m++
		dia[m] = i
	}

	if m&1 != 0 {
		p := dia[m/2+1]
		dp(p, 0, 0)
		fmt.Println(mul(I2, f[p][1][1]))
	} else {
		p := dia[m/2]
		q := dia[m/2+1]
		dp(p, q, 1)
		dp(q, p, 1)
		fmt.Println(mul((f[p][1][0]+f[p][1][1]+f[p][1][2])%MOD, (f[q][1][0]+f[q][1][1]+f[q][1][2])%MOD))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
