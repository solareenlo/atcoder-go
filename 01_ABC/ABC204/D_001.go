package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
		sum += t[i]
	}

	dp := make([]float64, sum/2+1)
	for i := 0; i < n; i++ {
		for j := sum / 2; j >= t[i]; j-- {
			dp[j] = math.Max(dp[j], dp[j-t[i]]+float64(t[i]))
		}
	}
	fmt.Println(sum - int(dp[sum/2]))
}
