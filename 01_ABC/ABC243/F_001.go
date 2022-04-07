package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, p int
	fmt.Fscan(in, &n, &m, &p)

	initMod()

	w := make([]int, size)
	s := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &w[i])
		s += w[i]
	}

	dp := [size][size][size]int{}
	dp[0][0][0] = 1
	for i := 1; i <= n; i++ {
		w[i] = w[i] * powMod(s, mod-2) % mod
		for j := 0; j <= m; j++ {
			for k := 0; k <= p; k++ {
				dp[i][j][k] += dp[i-1][j][k]
				dp[i][j][k] %= mod
				if j < m {
					for l, W := 1, w[i]; l+k <= p; l, W = l+1, W*w[i]%mod {
						dp[i][j+1][k+l] += dp[i-1][j][k] * W % mod * invf[l] % mod
						dp[i][j+1][k+l] %= mod
					}
				}
			}
		}
	}

	fmt.Println(fact[p] * dp[n][m][p] % mod)
}

const mod = 998244353
const size = 120

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
