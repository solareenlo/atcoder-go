package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	dp := [1001][1001]int{}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			dp[i+1][j] += dp[i][j] * a[i]
			dp[i+1][j] %= mod
			if k > j {
				dp[i+1][j+1] += (dp[i][j] * (k - j)) % mod
				dp[i+1][j+1] %= mod
			}
		}
	}

	res := 0
	invN := invMod(n)
	fac := 1
	for j := 0; j <= n && j <= k; j++ {
		res += (dp[n][j] * fac) % mod
		res %= mod
		fac *= invN
		fac %= mod
	}
	fmt.Println(res)
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}
