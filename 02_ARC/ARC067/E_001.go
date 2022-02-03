package main

import "fmt"

func main() {
	var n, A, B, C, D int
	fmt.Scan(&n, &A, &B, &C, &D)

	dp := [N][N]int{}
	dp[0][A-1] = 1
	initMod()
	for i := 0; i <= n; i++ {
		for j := A; j <= B; j++ {
			dp[i][j] = dp[i][j-1]
			p := powMod(inv[j], C)
			for k := C; k <= D; k++ {
				if i-j*k < 0 {
					break
				}
				dp[i][j] += dp[i-j*k][j-1] * p % mod * inv[k]
				dp[i][j] %= mod
				p *= inv[j]
				p %= mod
			}
		}
	}
	fmt.Println(dp[n][B] * fac[n] % mod)
}

const mod = 1000000007
const N = 1005

var fac, inv [N]int

func initMod() {
	fac[0] = 1
	inv[0] = 1
	for i := int(1); i < N; i++ {
		fac[i] = (fac[i-1] * i) % mod
		inv[i] = invMod(fac[i])
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
