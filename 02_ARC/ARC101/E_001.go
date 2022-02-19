package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 6005
const mod = 1_000_000_007

var (
	dp  = [N][N]int{}
	g   = [N]int{}
	siz = [N]int{}
	f1  = [N]int{}
	n   int
	G   = make([][]int, N)
)

func dfs(u, bef int) {
	siz[u] = 1
	dp[u][siz[u]] = 1
	for _, v := range G[u] {
		if v != bef {
			dfs(v, u)
			for i := 1; i <= siz[u]; i++ {
				for j := 1; j <= siz[v]; j++ {
					f1[i+j] += dp[u][i] * dp[v][j] % mod
					f1[i+j] %= mod
				}
			}
			siz[u] += siz[v]
			for i := 1; i <= siz[u]; i++ {
				dp[u][i] = (f1[i] - dp[u][i]*g[v]%mod + mod) % mod
				f1[i] = 0
			}
		}
	}
	fc := 1
	for i := 2; i <= siz[u]; i += 2 {
		g[u] += dp[u][i] * fc % mod
		g[u] %= mod
		fc = fc * (i + 1) % mod
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	dfs(1, 0)
	fmt.Println(g[1])
}
