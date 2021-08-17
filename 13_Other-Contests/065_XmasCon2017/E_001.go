package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	dp := make([][]bool, 1001)
	for i := range dp {
		dp[i] = make([]bool, 1001)
	}
	dp[0][0] = true

	n := len(s)
	m := len(t)
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if dp[i][j] {
				if i < n && j < m && s[i] == t[j] {
					dp[i+1][j+1] = true
				}
				if i < n && s[i] == 'A' {
					dp[i+1][j] = true
				}
				if j < m && t[j] == 'B' {
					dp[i][j+1] = true
				}
			}
		}
	}
	if dp[n][m] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
