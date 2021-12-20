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

	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &b[i])
	}

	dp := make([][2][2]int, n+1)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = 1 << 60
			}
		}
	}
	dp[1][0][0] = a[1]
	dp[1][1][1] = 0
	for i := 2; i < n+1; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for prej := 0; prej < 2; prej++ {
					tmp0, tmp1 := 0, 0
					if j == 0 {
						tmp0 = a[i]
					}
					if j == prej {
						tmp1 = b[i-1]
					}
					dp[i][j][k] = min(dp[i][j][k], dp[i-1][prej][k]+tmp0+tmp1)
				}
			}
		}
	}

	res := 1 << 60
	for j := 0; j < 2; j++ {
		for k := 0; k < 2; k++ {
			tmp := 0
			if j == k {
				tmp = b[n]
			}
			res = min(res, dp[n][j][k]+tmp)
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
