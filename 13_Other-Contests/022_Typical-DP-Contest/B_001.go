package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	a := make([]int, 1001)
	for i := A; i > 0; i-- {
		fmt.Scan(&a[i])
	}
	b := make([]int, 1001)
	for i := B; i > 0; i-- {
		fmt.Scan(&b[i])
	}

	dp := [1001][1001]int{}
	for i := 1; i <= max(A, B); i++ {
		dp[0][i] = dp[0][i-1] + b[i]*((A+B+i+1)%2)
		dp[i][0] = dp[i-1][0] + a[i]*((A+B+i+1)%2)
	}

	for i := 1; i <= A; i++ {
		for j := 1; j <= B; j++ {
			if (A+B+i+j)%2 == 0 {
				dp[i][j] = max(dp[i-1][j]+a[i], dp[i][j-1]+b[j])
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp[A][B])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
