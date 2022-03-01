package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

var (
	dp = [2000][2020]int{}
	G  = make([][]int, 2000)
)

func dfs(u int) int {
	ch := 0
	dp[u][0] = 1
	for _, v := range G[u] {
		c := dfs(v)
		tmp := make([]int, ch+c+1)
		for l := 0; l <= ch; l++ {
			for r := 0; r <= c; r++ {
				tmp[l+r] += dp[u][l] * dp[v][r]
				tmp[l+r] %= mod
			}
		}
		ch += c
		for i := 0; i <= ch; i++ {
			dp[u][i] = tmp[i]
		}
	}
	for i := ch; i >= 0; i-- {
		dp[u][i+1] += dp[u][i] * (ch - i) % mod
		dp[u][i+1] %= mod
	}
	ch++
	return ch
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var p int
		fmt.Fscan(in, &p)
		G[p-1] = append(G[p-1], i)
	}

	dfs(0)

	ans := 0
	fac := 1
	for i := n - 1; i >= 0; i-- {
		fac *= n - i
		fac %= mod
		now := fac * dp[0][i] % mod
		if i&1 != 0 {
			ans -= now
			ans += mod
			ans %= mod
		} else {
			ans += now
			ans %= mod
		}
	}
	fmt.Println(ans)
}
