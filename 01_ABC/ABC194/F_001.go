package main

import "fmt"

func to_int(c byte) int {
	if 47 < c && c < 58 {
		return int(c) - 48
	}
	return int(c) - 55
}

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)
	n := len(s)

	dp := [1 << 20][17]int{}
	m := map[int]struct{}{}
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		dp[i+1][0]++
		dp[i+1][0] %= mod
		for j := 0; j < 17; j++ {
			dp[i+1][j] += dp[i][j] * j
			dp[i+1][j] %= mod
			if j < 16 {
				if j == 0 {
					dp[i+1][j+1] += dp[i][j] * (16 - j - 1)
				} else {
					dp[i+1][j+1] += dp[i][j] * (16 - j)
				}
				dp[i+1][j+1] %= mod
			}
		}
		j := 0
		if i == 0 {
			j = 1
		}
		for ; j < to_int(s[i]); j++ {
			cnt := 1
			if _, ok := m[j]; ok {
				cnt = 0
			}
			dp[i+1][len(m)+cnt]++
			dp[i+1][len(m)+cnt] %= mod
		}
		m[to_int(s[i])] = struct{}{}
	}

	if len(m) == k {
		fmt.Println((dp[n][k] + 1) % mod)
	} else {
		fmt.Println(dp[n][k] % mod)
	}
}
