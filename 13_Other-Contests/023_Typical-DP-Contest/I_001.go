package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var S string
	fmt.Fscan(in, &S)
	N := len(S)

	var dp [310][310]int
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			dp[i][j] = 0
		}
	}

	for W := 3; W <= N; W++ {
		for l := 0; l+W <= N; l++ {
			r := l + W - 1
			if dp[l+2][r-1] == W-3 && S[l] == 'i' && S[l+1] == 'w' && S[r] == 'i' {
				dp[l][r] = W
				continue
			} else if dp[l+1][r-2] == W-3 && S[l] == 'i' && S[r-1] == 'w' && S[r] == 'i' {
				dp[l][r] = W
				continue
			}

			for m := l; m < r; m++ {
				dp[l][r] = max(dp[l][r], dp[l][m]+dp[m+1][r])
			}
		}
	}

	fmt.Println(dp[0][N-1] / 3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
