package main

import "fmt"

func main() {
	const MOD = 998244353
	const N = 10000010
	var fac, inv, f [N]int
	var g [9090]int

	var C func(int, int) int
	C = func(n, m int) int {
		return fac[n] * inv[n-m] % MOD * inv[m] % MOD
	}

	var n, m, K int
	fmt.Scan(&n, &m, &K)
	if n <= 3 {
		fmt.Println(1)
		return
	}
	if (K & 1) != 0 {
		K--
	}
	fac[0], fac[1], inv[0], inv[1] = 1, 1, 1, 1
	for i := 2; i <= K+n; i++ {
		fac[i] = fac[i-1] * i % MOD
	}
	for i := 2; i <= K+n; i++ {
		inv[i] = (MOD - MOD/i) * inv[MOD%i] % MOD
	}
	for i := 2; i <= K+n; i++ {
		inv[i] = inv[i-1] * inv[i] % MOD
	}
	for i := 0; i <= K; i++ {
		f[i] = C(i+n-1, n-1)
	}
	for i := 2; i <= K; i++ {
		f[i] = (f[i] + f[i-2]) % MOD
	}

	ans := 0
	for i := 0; i <= n && i*(m+1) <= K; i++ {
		var tmp int
		if (i & 1) != 0 {
			tmp = MOD - C(n, i)
		} else {
			tmp = C(n, i)
		}
		ans = (ans + tmp*f[K-i*(m+1)]%MOD) % MOD
	}
	for i := 0; i <= 2*m && i <= K; i++ {
		for j := 0; j <= n-2 && j*(m+1) <= i; j++ {
			var tmp int
			if (j & 1) != 0 {
				tmp = MOD - C(n-2, j)
			} else {
				tmp = C(n-2, j)
			}
			g[i] = (g[i] + tmp*C(i-j*(m+1)+n-3, n-3)%MOD) % MOD
		}
	}
	for i := 2; i <= 2*m && i <= K; i++ {
		g[i] = (g[i] + g[i-2]) % MOD
	}
	for i := 2; i <= 2*m && i <= K; i++ {
		ans = (ans + ((MOD-n)*g[min(i-2, K-i)]%MOD)*min(i+1, 2*m-i+1)%MOD) % MOD
	}
	for i := range g {
		g[i] = 0
	}
	for i := 0; i <= 3*m && i <= K; i++ {
		for j := 0; j <= n-3 && j*(m+1) <= i; j++ {
			var tmp int
			if (j & 1) != 0 {
				tmp = MOD - C(n-3, j)
			} else {
				tmp = C(n-3, j)
			}
			g[i] = (g[i] + tmp*C(i-j*(m+1)+n-4, n-4)%MOD) % MOD
		}
	}
	for i := 2; i <= 3*m && i <= K; i++ {
		g[i] = (g[i] + g[i-2]) % MOD
	}
	for i := 2; i <= 3*m && i <= K; i++ {
		for j := max(i-m+1, 0) / 2; j <= m && 2*j+2 <= i; j++ {
			ans = (ans + (n*g[min(i-2*j-2, K-i)]%MOD)*(2*(j-max(i-j-m, 0))+1)%MOD) % MOD
		}
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
