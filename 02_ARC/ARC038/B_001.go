package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([]string, h)
	for i := range s {
		fmt.Scan(&s[i])
	}

	dp := [101][101]bool{}
	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			if s[i][j] == '.' && !dp[i+1][j] && !dp[i+1][j+1] && !dp[i][j+1] {
				dp[i][j] = true
			}
		}
	}

	if dp[0][0] {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
