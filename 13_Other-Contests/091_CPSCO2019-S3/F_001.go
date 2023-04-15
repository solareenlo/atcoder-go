package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var dp [301][301][301]int
	dp[0][0][0] = 1
	for i := 0; i < a+b; i++ {
		for j := 0; j < a+b; j++ {
			for k := 0; k < a+b; k++ {
				v := dp[i][j][k]
				dp[i+1][j][k+1] = (dp[i+1][j][k+1] + v) % mod
				dp[i+1][j+1][k] = (dp[i+1][j+1][k] + v*k%mod) % mod
				dp[i+1][j][k] = (dp[i+1][j][k] + v*k%mod) % mod
				if k != 0 {
					dp[i+1][j+1][k-1] = (dp[i+1][j+1][k-1] + (v*k%mod)*k%mod) % mod
				}
			}
		}
	}
	ans := dp[a+b][a][0]
	for i := 0; i < a+b; i++ {
		ans = divMod(ans*(n-i)%mod, (i + 1))
	}
	fmt.Println(ans)
}

const mod = 1000000007

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
