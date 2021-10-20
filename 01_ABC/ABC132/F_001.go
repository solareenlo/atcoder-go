package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	num := make([]int, 0)
	for i := 1; i*i < n+1; i++ {
		num = append(num, i)
		if i != n/i {
			num = append(num, n/i)
		}
	}
	sort.Ints(num)

	mod := int(1e9 + 7)
	m := len(num)
	dp := make([]int, m)
	dp[0] = 1
	for i := 0; i < k; i++ {
		for i, j := 0, len(dp)-1; i < j; i, j = i+1, j-1 {
			dp[i], dp[j] = dp[j], dp[i]
		}
		for j := m - 1; j >= 1; j-- {
			dp[j-1] += dp[j]
			dp[j-1] %= mod
			dp[j] *= num[j] - num[j-1]
			dp[j] %= mod
		}
	}

	sum := 0
	for i := range dp {
		sum += dp[i]
		sum %= mod
	}
	fmt.Println(sum)
}
