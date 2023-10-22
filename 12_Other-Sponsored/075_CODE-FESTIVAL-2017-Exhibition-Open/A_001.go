package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAXN = 2005
	const MOD = 1000000007

	var dp [MAXN][MAXN][2]int
	var fact [MAXN]int

	fact[0] = 1
	for i := 1; i < MAXN; i++ {
		fact[i] = fact[i-1] * i % MOD
	}

	var n int
	fmt.Fscan(in, &n)
	var p, sz [MAXN]int
	var gph [MAXN][]int
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		gph[p[i]] = append(gph[p[i]], i)
	}
	for i := n; i > 0; i-- {
		var aux [MAXN][3]int
		aux[0][0] = 1
		for _, j := range gph[i] {
			var nxt [MAXN][3]int
			for k := 0; k <= sz[j]; k++ {
				for l := 0; l <= sz[i]; l++ {
					for m := 0; m < 3; m++ {
						nxt[k+l][m] += dp[j][k][0] * aux[l][m] % MOD
						if m < 2 {
							nxt[k+l+1][m+1] += dp[j][k][1] * aux[l][m] % MOD
						}
					}
				}
			}
			sz[i] += sz[j]
			for j := 0; j <= sz[i]; j++ {
				for k := 0; k < 3; k++ {
					aux[j][k] = nxt[j][k] % MOD
				}
			}
		}
		for j := 0; j <= sz[i]; j++ {
			dp[i][j][0] = aux[j][0] + aux[j][1]*2 + aux[j][2]*2
			dp[i][j][1] = aux[j][0] + aux[j][1]
			dp[i][j][0] %= MOD
			dp[i][j][1] %= MOD
		}
		sz[i]++
	}
	ret := 0
	for i := 0; i <= n; i++ {
		cur := dp[1][i][0] * fact[n-i] % MOD
		if (i & 1) != 0 {
			cur = MOD - cur
		}
		ret += cur
	}
	fmt.Println(ret % MOD)
}
