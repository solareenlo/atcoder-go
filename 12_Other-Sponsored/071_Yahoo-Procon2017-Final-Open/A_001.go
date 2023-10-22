package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var dp [100010][5]int
	for i := 0; i < 5; i++ {
		dp[0][i] = i
	}
	for i := 0; i <= len(s); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= 10; j++ {
			tmp := 0
			if s[i-1] != "yahoo"[(j-1)%5] {
				tmp = 1
			}
			dp[i][j%5] = min(dp[i-1][(j-1)%5]+tmp, min(dp[i-1][j%5], dp[i][(j-1)%5])+1)
		}
	}
	fmt.Println(dp[len(s)][0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
