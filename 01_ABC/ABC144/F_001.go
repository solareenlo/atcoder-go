package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		g[a] = append(g[a], b)
	}
	res := 1e9
	for u := 0; u < n; u++ {
		dp := make([]float64, n)
		for i := n - 1; i >= 0; i-- {
			sum, maxi := 0.0, 0.0
			cnt := 0
			for _, j := range g[i] {
				cnt++
				sum += dp[j]
				maxi = max(maxi, dp[j])
			}
			if cnt == 0 {
				dp[i] = 0.0
				continue
			}
			if u == i {
				if cnt > 1 {
					cnt--
					sum -= maxi
				}
			}
			dp[i] = sum/float64(cnt) + 1
		}
		res = min(res, dp[0])
	}
	fmt.Printf("%.10f\n", res)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
