package main

import "fmt"

func main() {
	var n, a, b, m int
	fmt.Scan(&n, &a, &b, &m)
	a--
	b--

	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, 0)
	}
	var x, y int
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		x--
		y--
		d[x] = append(d[x], y)
		d[y] = append(d[y], x)
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][a] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				continue
			}
			if j == b {
				fmt.Println(dp[i][j])
				return
			}
			for k := 0; k < len(d[j]); k++ {
				to := d[j][k]
				dp[i+1][to] = (dp[i+1][to] + dp[i][j]) % int(1e9+7)
			}
		}
	}
}
