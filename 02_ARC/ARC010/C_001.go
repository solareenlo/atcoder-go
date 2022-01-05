package main

import "fmt"

func main() {
	var n, m, y, z int
	fmt.Scan(&n, &m, &y, &z)

	const N = 2222
	dp := [N][N]int{}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1 << 60
		}
	}

	id := make([]int, N)
	x := make([]int, N)
	for i := 0; i < m; i++ {
		var s string
		fmt.Scan(&s, &x[i])
		dp[0][i] = 0
		id[s[0]] = i
	}

	var s string
	fmt.Scan(&s)
	for i := 0; i < n; i++ {
		for j := (1 << m) - 1; j >= 0; j-- {
			t := dp[0][1<<m]
			p := id[s[i]]
			for k := 0; k < m; k++ {
				tmp := 0
				if j != 0 && k == p {
					tmp = 1
				}
				t = max(t, dp[j][k]+x[p]+tmp*y)
			}
			dp[j|1<<p][p] = max(dp[j|1<<p][p], t)
		}
	}

	res := 0
	for i := 0; i < 1<<m; i++ {
		for j := 0; j < m; j++ {
			tmp := 0
			if i+1 == 1<<m {
				tmp = 1
			}
			res = max(res, dp[i][j]+tmp*z)
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
