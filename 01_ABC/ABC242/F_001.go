package main

import "fmt"

func main() {
	var n, m, b, w int
	fmt.Scan(&n, &m, &b, &w)

	initMod()

	const N = 60
	dp := [N][N]int{}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = nCrMod(i*j, b)
			for x := 1; x <= i; x++ {
				for y := 1; y <= j; y++ {
					if x != i || y != j {
						dp[i][j] = dp[i][j] - dp[x][y]*nCrMod(i, x)%mod*nCrMod(j, y)%mod + mod
						dp[i][j] %= mod
					}
				}
			}
			ans += dp[i][j] * nCrMod((n-i)*(m-j), w) % mod * nCrMod(n, i) % mod * nCrMod(m, j) % mod
			ans %= mod
		}
	}
	fmt.Println(ans)
}

const mod = 998244353
const size = 3000

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
