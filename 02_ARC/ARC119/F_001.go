package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	var s string
	fmt.Fscan(in, &n, &k, &s)
	s = " " + s

	const N = 4040
	dp := [N][2][N][4]int{}
	if s[1] != 'B' {
		dp[1][0][0][2] = 1
	}
	if s[1] != 'A' {
		dp[1][1][0][2] = 1
	}
	mod := 1_000_000_007
	ans := 0
	for i := 1; i < n; i++ {
		for j := 0; j < 2; j++ {
			if int(s[i]) != 'A'+(j^1) {
				for t := 0; t <= k; t++ {
					for l := -1; l <= 2; l++ {
						x := dp[i][j][t][l+1]
						if i < n-1 {
							dp[i+1][j][t][min(l+1, 2)+1] += x
							dp[i+1][j][t][min(l+1, 2)+1] %= mod
							if t+l >= 0 {
								dp[i+1][j^1][t+l][min(t+1, t+l+1)-t-l+1] += x
								dp[i+1][j^1][t+l][min(t+1, t+l+1)-t-l+1] %= mod
							}
						} else if min(t+1, t+l+1) <= k {
							ans += x
							ans %= mod
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
