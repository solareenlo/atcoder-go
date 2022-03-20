package main

import "fmt"

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	dp := make([]map[int]float64, 110)
	for i := range dp {
		dp[i] = make(map[int]float64)
	}

	dp[0][1] = 1.0
	for i := 0; i < n; i++ {
		for g, p := range dp[i] {
			for k := 1; k <= 6; k++ {
				dp[i+1][gcd(g*k, d)] += p / 6
			}
		}
	}

	fmt.Println(dp[n][d])
}
