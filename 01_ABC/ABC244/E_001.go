package main

import "fmt"

func main() {
	var n, m, k, s, t, x int
	fmt.Scan(&n, &m, &k, &s, &t, &x)
	s--
	t--
	x--

	type pair struct{ u, v int }
	edge := make([]pair, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&edge[i].u, &edge[i].v)
		edge[i].u--
		edge[i].v--
	}

	const mod = 998244353
	dp := [2001][2001][2]int{}
	dp[0][s][0] = 1
	for i := 0; i < k; i++ {
		for _, e := range edge {
			u := e.u
			v := e.v
			for j := 0; j < 2; j++ {
				tmp := 0
				if v == x {
					tmp = 1
				}
				dp[i+1][v][j^tmp] += dp[i][u][j]
				dp[i+1][v][j^tmp] %= mod
				tmp = 0
				if u == x {
					tmp = 1
				}
				dp[i+1][u][j^tmp] += dp[i][v][j]
				dp[i+1][u][j^tmp] %= mod
			}
		}
	}
	fmt.Println(dp[k][t][0])
}
