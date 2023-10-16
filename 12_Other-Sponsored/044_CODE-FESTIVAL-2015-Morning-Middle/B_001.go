package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	maxi := 0
	for k := 0; k < n+1; k++ {
		l := s[0:k]
		r := s[k:]
		L := len(l)
		R := len(r)
		var dp [102][102]int
		for i := 0; i < L; i++ {
			for j := 0; j < R; j++ {
				if l[i] == r[j] {
					dp[i+1][j+1] = dp[i][j] + 1
				} else {
					dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
				}
			}
		}
		maxi = max(maxi, dp[L][R])
	}
	fmt.Println(n - maxi*2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
