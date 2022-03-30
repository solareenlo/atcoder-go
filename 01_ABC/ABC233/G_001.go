package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	s := make([]string, n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	dp := [50][50][50][50]int{}
	for x1 := n - 1; x1 >= 0; x1-- {
		for y1 := n - 1; y1 >= 0; y1-- {
			for x2 := x1; x2 <= n-1; x2++ {
				for y2 := y1; y2 <= n-1; y2++ {
					if x1 == x2 && y1 == y2 {
						if s[x1][y1] == '#' {
							dp[x1][x2][y1][y2] = 1
						}
					} else {
						dp[x1][x2][y1][y2] = max(x2-x1+1, y2-y1+1)
					}
					for k := x1; k < x2; k++ {
						dp[x1][x2][y1][y2] = min(dp[x1][x2][y1][y2], dp[x1][k][y1][y2]+dp[k+1][x2][y1][y2])
					}
					for k := y1; k < y2; k++ {
						dp[x1][x2][y1][y2] = min(dp[x1][x2][y1][y2], dp[x1][x2][y1][k]+dp[x1][x2][k+1][y2])
					}
				}
			}
		}
	}

	fmt.Println(dp[0][n-1][0][n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
