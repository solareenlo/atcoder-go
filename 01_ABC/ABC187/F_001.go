package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	g := [20]int{}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		g[a] |= 1 << b
		g[b] |= 1 << a
	}

	dp := make([]int, 1<<18)
	for i := 0; i < 1<<18; i++ {
		dp[i] = 1 << 60
	}
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if dp[j] == 1 && (j&g[i]) == j {
				dp[j|1<<i] = 1
			}
		}
	}

	for i := 0; i < 1<<n; i++ {
		for j := i - 1; j > 1/2; j-- {
			j &= i
			if j == 0 {
				break
			}
			if dp[j]+dp[i^j] < dp[i] {
				dp[i] = dp[j] + dp[i^j]
			}
		}
	}

	fmt.Println(dp[1<<n-1])
}
