package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	A := make([]int, n)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
		sum[i+1] = sum[i] + A[i]
	}

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	dp[1][0]++

	mod := int(1e9 + 7)
	res := 0
	for i := 0; i < n; i++ {
		for j := n; 1 < j+1; j-- {
			dp[j+1][sum[i+1]%(j+1)] += dp[j][sum[i+1]%j]
			dp[j+1][sum[i+1]%(j+1)] %= mod
			if i == n-1 {
				res += dp[j][sum[i+1]%j]
				res %= mod
			}
		}
	}
	fmt.Println(res)
}
