package main

import "fmt"

func main() {
	var dp [1001][1001][11]int

	var s, t string
	fmt.Scan(&s, &t)
	n := len(s)
	m := len(t)

	var k int
	fmt.Scan(&k)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for h := 0; h <= k; h++ {
				tmp := 0
				if h != 0 {
					tmp = dp[i][j][h-1] + 1
				}
				tmp1 := 0
				if s[i] == t[j] {
					tmp1 = 1
				}
				dp[i+1][j+1][h] = max(dp[i][j+1][h], dp[i+1][j][h], dp[i][j][h]+tmp1, tmp)
			}
		}
	}

	r := 0
	for h := 0; h <= k; h++ {
		r = max(r, dp[n][m][h])
	}
	fmt.Println(r)
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
