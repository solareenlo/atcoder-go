package main

import "fmt"

func main() {
	const N = 10001

	var a, b string
	fmt.Scan(&a, &b)
	n := len(a)

	initMod()

	s := make([]int, N)
	for i := 0; i < n; i++ {
		if a[i]^48 != 0 {
			s[b[i]^48]++
		}
	}

	dp := [N][N]int{}
	for i := 0; i <= s[0]; i++ {
		dp[0][i] = fact[i] * fact[i] % mod
	}
	for i := 1; i <= s[1]; i++ {
		for j := 1; j <= s[0]; j++ {
			dp[i][j] = dp[i-1][j]*i*j%mod + dp[i][j-1]*j*j%mod
			dp[i][j] %= mod
		}
	}

	ans := 0
	for i := 0; i <= s[1]; i++ {
		ans += dp[s[1]-i][s[0]] * fact[s[1]] % mod * fact[s[0]+s[1]] % mod * invf[s[1]-i] % mod * invf[s[1]+s[0]-i] % mod
		ans %= mod
	}
	fmt.Println(ans)
}

const mod = 998244353
const size = 10010

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
