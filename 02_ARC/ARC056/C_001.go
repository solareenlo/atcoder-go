package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	t := 0
	w := [20][20]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&w[i][j])
			t += w[i][j]
		}
	}

	dp := make([]int, 1<<n)
	for bit := 0; bit < 1<<n; bit++ {
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				if (bit>>i)&(bit>>j)&1 != 0 {
					dp[bit] += w[i][j]
				}
			}
		}
		if bit != 0 {
			dp[bit] += k
		}
		for i := bit; i > 0; i = (i - 1) & bit {
			dp[bit] = max(dp[bit], dp[bit^i]+dp[i])
		}
	}

	fmt.Println(dp[(1<<n)-1] - t/2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
