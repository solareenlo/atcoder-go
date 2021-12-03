package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	cnt := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		cnt[i] = 1
	}

	mod := 998244353
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		for j := 0; j < 1<<n; j++ {
			if (j>>a&1 != 0) && (j>>b&1 != 0) {
				cnt[j] *= 2
				cnt[j] %= mod
			}
		}
	}

	dp := make([]int, 1<<17)
	res := [17]int{}
	for i := 1; i < 1<<n; i += 2 {
		now := cnt[i]
		for j := i - 1; j >= 0; j-- {
			j &= i
			now = (now - (dp[j] * cnt[i^j] % mod) + mod) % mod
		}
		dp[i] = now
		for j := 1; j < n; j++ {
			if i>>j&1 != 0 {
				res[j] += now * cnt[(1<<n)-1^i] % mod
				res[j] %= mod
			}
		}
	}

	for i := 1; i < n; i++ {
		fmt.Println(res[i])
	}
}
