package main

import "fmt"

func main() {
	var s string

	const N = 145
	const P = 998244353

	var f [N][N]int

	fmt.Scan(&s)
	n := len(s)
	s = " " + s
	for i := 0; i < n; i++ {
		for j := 97; j <= 122; j++ {
			f[i][j] = n + 1
		}
		for j := n; j > i; j-- {
			f[i][s[j]] = j
		}
	}
	ans := 0
	for p := 1; p < n; p++ {
		var dp [N][N]int
		dp[0][p] = 1
		for i := 0; i < p; i++ {
			for j := p; j < n; j++ {
				for t := 97; t <= 122; t++ {
					dp[f[i][t]][f[j][t]] = (dp[f[i][t]][f[j][t]] + dp[i][j]) % P
				}
			}
		}
		for i := n; i > p; i-- {
			ans = (ans + dp[p][i]) % P
		}
	}
	fmt.Println(ans)
}
