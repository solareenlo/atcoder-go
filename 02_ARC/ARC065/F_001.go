package main

import (
	"fmt"
)

func main() {
	var n, m int
	var s string
	fmt.Scan(&n, &m, &s)

	s += "0"

	n++
	p := make([]int, n+1)
	o := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = p[i-1] + int(s[i-1]) - 48
		o[i] = i
	}

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		if o[l] < r {
			o[l] = r
		}
	}

	for i := 1; i <= n; i++ {
		if o[i] <= o[i-1] {
			o[i] = o[i-1]
		}
	}

	dp := [3005][3005]int{}
	mod := int(1e9 + 7)
	dp[1][p[o[1]]] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= o[i-1]-i+1; j++ {
			dp[i][j+p[o[i]]-p[o[i-1]]] = (dp[i-1][j] + dp[i-1][j+1]) % mod
		}
	}
	fmt.Println(dp[n][0])
}
