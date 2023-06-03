package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	t := ""
	for i := 0; i < n; i++ {
		t += string(s[n-1-i])
	}
	var dp [2009][2009]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	fmt.Println(dp[n][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
