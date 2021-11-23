package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	x := make([]int, m+1)
	y := make([]int, m+1)
	z := make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Scan(&x[i], &y[i], &z[i])
	}

	N := int(1 << n)
	dp := make([]int, N)
	dp[0] = 1
	sf := make([]int, n+1)
	pc := make([]int, N)
	for i := 1; i < N; i++ {
		var j int
		pc[i] = pc[i>>1] + (i & 1)
		sf[n] = 0
		for j = n - 1; j >= 0; j-- {
			sf[j] = sf[j+1] + (i >> j & 1)
		}
		for j = 1; j <= m; j++ {
			if x[j] == pc[i] && pc[i]-sf[y[j]] > z[j] {
				break
			}
		}
		if j > m {
			for j := 0; j < n; j++ {
				if i>>j&1 != 0 {
					dp[i] += dp[i^(1<<j)]
				}
			}
		}
	}

	fmt.Println(dp[N-1])
}
