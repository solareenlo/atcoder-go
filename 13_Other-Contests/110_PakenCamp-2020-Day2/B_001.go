package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)
	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &s[i])
	}
	var dp [2002][2002]int
	for i := 0; i < h+1; i++ {
		for j := 0; j < w+1; j++ {
			dp[i][j] = int(1e9)
		}
	}
	dp[0][0] = 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != 'S' {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
			} else {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j])
			}
			if s[i][j] != 'E' {
				dp[i][j+1] = min(dp[i][j+1], dp[i][j]+1)
			} else {
				dp[i][j+1] = min(dp[i][j+1], dp[i][j])
			}
		}
	}
	fmt.Println(dp[h-1][w])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
