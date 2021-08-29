package main

import "fmt"

func main() {
	var n, m, d, a int
	fmt.Scan(&n, &m, &d)

	dp := [30][100001]int{}
	for i := 0; i < n; i++ {
		dp[0][i] = i
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		dp[0][a-1], dp[0][a] = dp[0][a], dp[0][a-1]
	}

	for i := 1; i < 30; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	res := [100001]int{}
	for i := 0; i < n; i++ {
		now := i
		for j := 0; j < 30; j++ {
			if d>>j&1 == 1 {
				now = dp[j][now]
			}
		}
		res[now] = i
	}
	for i := 0; i < n; i++ {
		fmt.Println(res[i] + 1)
	}
}
