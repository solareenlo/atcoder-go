package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	dp := make([]int, m+1)
	fact := make([]int, n+1)
	inv := make([]int, n+1)
	dp[0] = 1
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = i * fact[i-1] % mod
	}
	inv[n] = divMod(1, fact[n])
	for i := n - 1; i >= 0; i-- {
		inv[i] = (i + 1) * inv[i+1] % mod
	}

	for i := 2; i <= m; i += 2 {
		for j := 0; j <= n && j <= i; j += 2 {
			dp[i] += fact[n] * inv[n-j] % mod * inv[j] % mod * dp[(i-j)/2] % mod
			dp[i] %= mod
		}
	}
	fmt.Println(dp[m])
}

const mod = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
