package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MAXN = 30
	const MAXV = 20
	const MAXB = 600
	const MOD = 998244353

	var n, v int
	fmt.Fscan(in, &n, &v)

	var f [2][MAXN + 2][MAXV + 2][MAXV + 2][MAXN*(MAXV+1) + MAXB + 2]int
	var pow_v [MAXN + 2]int
	var ans [MAXN + 5]int

	pow_v[0] = 1
	for i := 1; i <= n; i++ {
		pow_v[i] = pow_v[i-1] * v % MOD
	}
	b := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}
	tmp := b[1 : n+1]
	sort.Ints(tmp)
	cur := 0
	var nxt int
	f[cur][0][0][v+1][0] = 1
	for i := 1; i <= n; i++ {
		nxt = cur ^ 1
		for j := 0; j <= i; j++ {
			for k := range f[nxt][j] {
				for l := range f[nxt][j][k] {
					for m := range f[nxt][j][k][l] {
						f[nxt][j][k][l][m] = 0
					}
				}
			}
		}
		for j := 0; j <= i; j++ {
			for mx := 0; mx <= v; mx++ {
				for mn := mx; mn <= v+1; mn++ {
					for sum := 0; sum <= b[i]+(v+1)*i; sum++ {
						if f[cur][j][mx][mn][sum] == 0 {
							continue
						}
						if mx > 0 && sum <= b[i] {
							ans[j] = (ans[j] + f[cur][j][mx][mn][sum]*pow_v[n-i+1]%MOD*sum) % MOD
						}
						if mn == v+1 || sum+mn > b[i] {
							var n_mx, n_mn int
							n_sum := max(sum, b[i])
							n_mx = n_sum - max(n_sum-mx, b[i])
							if mn == v+1 {
								n_mn = v + 1
							} else {
								n_mn = mn + sum - n_sum
							}
							for t := 1; t < n_mx; t++ {
								f[nxt][j+1][n_mx][n_mn][n_sum+t] = (f[nxt][j+1][n_mx][n_mn][n_sum+t] + f[cur][j][mx][mn][sum]) % MOD
							}
							for t := max(1, n_mx); t < n_mn; t++ {
								f[nxt][j][n_mx][t][n_sum] = (f[nxt][j][n_mx][t][n_sum] + f[cur][j][mx][mn][sum]) % MOD
								f[nxt][j+1][t][n_mn][n_sum+t] = (f[nxt][j+1][t][n_mn][n_sum+t] + f[cur][j][mx][mn][sum]) % MOD
							}
							f[nxt][j][n_mx][n_mn][n_sum] = (f[nxt][j][n_mx][n_mn][n_sum] + f[cur][j][mx][mn][sum]*(v-n_mn+1)) % MOD
						}
					}
				}
			}
		}
		cur = nxt
	}
	for j := 1; j <= n; j++ {
		for mx := 1; mx <= v; mx++ {
			for mn := mx; mn <= v+1; mn++ {
				for sum := 0; sum <= b[n]+(v+1)*n; sum++ {
					if f[cur][j][mx][mn][sum] == 0 {
						continue
					}
					ans[j] = (ans[j] + f[cur][j][mx][mn][sum]*sum) % MOD
				}
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", ans[i])
	}
	fmt.Fprintln(out)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
