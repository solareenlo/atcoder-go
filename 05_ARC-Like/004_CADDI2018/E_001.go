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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	dp := [2][200005][16]int{}
	for k := 0; k < 2; k++ {
		for i := 0; i < 16; i++ {
			dp[k][n-1][i] = i * 2
		}
		for i := n - 1; i > 0; i-- {
			t := 0
			if a[i-1] <= a[i] {
				for ; a[i-1] <= a[i]>>(t*2+2); t++ {
				}
			} else {
				for ; a[i-1] > a[i]<<(-t*2); t-- {
				}
			}
			for j := 0; j < 16; j++ {
				if j-t > 15 {
					dp[k][i-1][j] = j*2 + (dp[k][i][15] + (j-t-15)*(n-i)*2)
				} else {
					dp[k][i-1][j] = j*2 + dp[k][i][max(0, j-t)]
				}
			}
		}
		if k == 0 {
			for i := 0; i < n/2; i++ {
				a[i], a[n-1-i] = a[n-1-i], a[i]
			}
		}
	}

	ans := 1 << 60
	for i := 0; i < n; i++ {
		ans = min(ans, dp[0][i][0]+dp[1][n-i][0]+i)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
