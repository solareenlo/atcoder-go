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
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
	}

	dp := [3005][3005]int{}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 1 << 60
		}
	}
	dp[0][1] = 0
	for i := 0; i < n; i++ {
		for j := 1; j <= n; j++ {
			tmp := -1
			if s[i] == ')' {
				tmp = 1
			}
			dp[i+1][j] = min(dp[i][j+tmp], dp[i][j]+d[i], dp[i][j-tmp]+c[i])
		}
	}
	fmt.Println(dp[n][1])
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
