package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 404 * 2
	l := make([]int, N)
	r := make([]int, N)
	s := [N][N]int{}
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		s[l[i]][r[i]]++
	}

	for l := 1; l <= n; l++ {
		for i, j := 1, l; j <= n; i, j = i+1, j+1 {
			s[i][j] += s[i+1][j] + s[i][j-1] - s[i+1][j-1]
		}
	}

	dp := [N][N]int{}
	dp[0][0] = -1
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			for k := i + 1; k <= n; k++ {
				dp[k][j+s[i+1][k-1]] = (dp[k][j+s[i+1][k-1]] - dp[i][j] + mod) % mod
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if m-j-s[i+1][n] != 0 {
				ans += divMod(dp[i][j]*m%mod, (m-j-s[i+1][n]+mod)%mod)
				ans %= mod
			}
		}
	}

	fmt.Println(ans)
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
