package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

func ksm(a, b, c int) int {
	for b > 0 {
		if b&1 != 0 {
			c = c * a % mod
		}
		b /= 2
		a = a * a % mod
	}
	return c
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	fa := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &fa[i])
	}

	g := make([][]int, 200200)
	f := make([][]int, 200200)
	tag := make([]int, n+1)
	inv_2 := (mod + 1) / 2
	for i := n; i > 0; i-- {
		f[i] = append(f[i], inv_2)
		g[i] = append(g[i], inv_2)
		tag[i] = len(f[i])
		u := fa[i]
		if len(f[u]) < len(f[i]) {
			f[u], f[i] = f[i], f[u]
			g[u], g[i] = g[i], g[u]
			tag[u], tag[i] = tag[i], tag[u]
		}
		for j := 0; j < tag[i]; j++ {
			f[i][j] = (mod + 1 - g[i][j]) % mod
		}
		for j := 0; j < len(f[i]); j++ {
			nj := len(f[i]) - 1 - j
			nk := len(f[u]) - 1 - j
			if nk < tag[u] {
				tag[u]--
				f[u][nk] = (mod + 1 - g[u][nk]) % mod
			}
			g[u][nk] = (g[u][nk]*f[i][nj] + f[u][nk]*g[i][nj]) % mod
			f[u][nk] = f[u][nk] * f[i][nj] % mod
		}
	}

	ans := inv_2
	for _, i := range g[0] {
		ans = (ans + i) % mod
	}
	fmt.Println(ksm(2, n+1, ans))
}
