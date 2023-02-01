package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s int
	fmt.Fscan(in, &n, &s)

	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	var dp [101][10001]int
	var ans [101][10001]string
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := s; j >= min(a[i], b[i]); j-- {
			if j >= a[i] && dp[i-1][j-a[i]] != 0 {
				dp[i][j] = 1
				ans[i][j] = ans[i-1][j-a[i]] + "H"
			}
			if j >= b[i] && dp[i-1][j-b[i]] != 0 {
				dp[i][j] = 1
				ans[i][j] = ans[i-1][j-b[i]] + "T"
			}
		}
	}

	if dp[n][s] != 0 {
		fmt.Println("Yes")
		fmt.Println(ans[n][s])
	} else {
		fmt.Println("No")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
