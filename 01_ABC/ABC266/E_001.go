package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var dp [101]float64
	for i := 1; i <= n; i++ {
		for j := 1; j <= 6; j++ {
			dp[i] += math.Max(dp[i-1], float64(j)) / 6
		}
	}
	fmt.Println(dp[n])
}
