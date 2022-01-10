package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	p := [2222][2222]float64{}
	a := make([]float64, 2222)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x, &a[i])
		for j := 0; j < x; j++ {
			var y int
			fmt.Scan(&y)
			fmt.Scan(&p[i][y-1])
		}
	}

	dp := make([]float64, 2222)
	for bit := (1 << n) - 2; bit >= 0; bit-- {
		dp[bit] = 1e99
		for i := 0; i < m; i++ {
			t := 0.0
			q := 0.0
			for k := 0; k < n; k++ {
				if (bit>>k)&1 != 0 {
					q += p[i][k] / 100.0
				} else {
					t += p[i][k] / 100.0 * dp[bit|(1<<k)]
				}
			}
			if q < 1.0 {
				dp[bit] = min(dp[bit], (t+a[i])/(1.0-q))
			}
		}
	}

	fmt.Println(dp[0])
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
