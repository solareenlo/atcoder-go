package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Scan(&a[i])
		sum += a[i]
	}

	if k == sum {
		fmt.Println(1)
		return
	}

	const inf = 1 << 60
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	dp[1] = 1
	sum = a[0]
	for i := 1; i < n; i++ {
		for j := i; j >= 0; j-- {
			dp[j+1] = min(dp[j+1], (dp[j]*(sum+a[i]))/sum+1)
		}
		sum += a[i]
	}

	for i := n; i >= 0; i-- {
		if dp[i] <= k {
			fmt.Println(i)
			return
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
