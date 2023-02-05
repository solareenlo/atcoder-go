package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353
	const N = 3030

	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)
	var inv [N]int
	inv[1] = 1
	for i := 2; i <= 3000; i++ {
		inv[i] = mod - (mod/i)*inv[mod%i]%mod
	}
	G := make([][]int, N)
	for i := range G {
		G[i] = make([]int, 0)
	}
	var ru [N]int
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		ru[u]++
		ru[v]++
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	var c [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}
	var f [N][3]int
	f[1][0] = 1
	ans := 0
	for d > 0 {
		d--
		var g [N][3]int
		for i := 1; i <= n; i++ {
			for _, j := range G[i] {
				for p := 0; p < 3; p++ {
					g[j][p] += f[i][p] * inv[ru[i]] % mod
					g[j][p] %= mod
				}
			}
		}
		for i := 1; i <= n; i++ {
			if c[i] != 0 {
				ans += g[i][2]
				ans %= mod
			} else {
				g[i][2] += (2*g[i][1]%mod + g[i][0]) % mod
				g[i][2] %= mod
				g[i][1] += g[i][0]
				g[i][1] %= mod
			}
		}
		for i := 1; i <= n; i++ {
			for j := 0; j < 3; j++ {
				f[i][j] = g[i][j]
			}
		}
	}
	fmt.Println(ans)
}
