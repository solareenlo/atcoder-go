package main

import "fmt"

func main() {
	initMod()

	var n, k int
	fmt.Scan(&n, &k)
	n++

	dp := [5005][5005]int{}
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			dp[i+1][j] += dp[i][j] * j
			dp[i+1][j] %= mod
			dp[i+1][j+1] += dp[i][j]
			dp[i+1][j+1] %= mod
		}
	}

	ans := 0
	for i := 0; i < n+1; i++ {
		x := 0
		for j := 0; j < i+1; j++ {
			x += dp[j][k] * dp[i-j][k] % mod
			x %= mod
		}
		rem := n - i
		y := 0
		for j := 0; j < rem+1; j++ {
			y += dp[rem][j]
			y %= mod
		}
		ans += x * y % mod * nCrMod(n, i) % mod
		ans %= mod
	}
	fmt.Println(ans)
}

const mod = 1000000007
const size = 5005

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
