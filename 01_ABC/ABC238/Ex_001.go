package main

import "fmt"

func main() {
	initMod()

	var N int
	var dir string
	fmt.Scan(&N, &dir)
	dp := [303][303]int{}
	for i := 0; i < N; i++ {
		dp[i][i+1] = 1
	}
	dp[N-1][0] = 1
	cost := [303][303]int{}
	for l := 2; l <= N; l++ {
		for i, j := 0, l; i < N; i, j = i+1, j+1 {
			if j == N {
				j = 0
			}
			for k, a := i+1, 0; a <= l-2; k, a = k+1, a+1 {
				if k == N {
					k = 0
				}
				coeff := 0
				len := 0
				if dir[i] == 'R' {
					coeff++
					len += a + 1
				}
				if dir[j] == 'L' {
					coeff++
					len += l - a - 1
				}
				coeff = coeff * nCrMod(l-2, a) % mod
				len = len * nCrMod(l-2, a) % mod
				dp[i][j] += coeff * dp[i][k] % mod * dp[k][j] % mod
				dp[i][j] %= mod
				cost[i][j] +=
					coeff*((cost[i][k]*dp[k][j]+dp[i][k]*cost[k][j])%mod)%mod +
						len*dp[i][k]%mod*dp[k][j]%mod
				cost[i][j] %= mod
			}
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans += cost[i][i]
		ans %= mod
	}
	ans = ans * invf[N] % mod
	fmt.Println(ans)
}

const mod = 998244353
const size = 303

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
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

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
