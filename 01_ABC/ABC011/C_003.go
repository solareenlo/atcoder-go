package main

import "fmt"

func main() {
	var n, ng1, ng2, ng3 int
	fmt.Scan(&n, &ng1, &ng2, &ng3)

	if n == ng1 || n == ng2 || n == ng3 {
		fmt.Println("NO")
		return
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = int(1e9)
	}
	dp[n] = 0
	for i := n; i >= 0; i-- {
		if i == ng1 || i == ng2 || i == ng3 {
			continue
		}
		for j := 1; j <= 3; j++ {
			if i-j >= 0 {
				dp[i-j] = min(dp[i-j], dp[i]+1)
			}
		}
	}
	res := "NO"
	if dp[0] <= 100 {
		res = "YES"
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
