package main

import "fmt"

const (
	N   = 400
	M   = 100000
	mod = 1000000007
)

var ()

func main() {
	var n, m, k, ans int
	var X, Y, pw [N]int
	var s [M]int
	var f [3][M]int

	fmt.Scan(&n, &m)
	k = (1 << n) - 1
	for i := 1; i <= m; i++ {
		fmt.Scan(&X[i], &Y[i])
	}
	pw[0] = 1
	for i := 1; i <= m; i++ {
		pw[i] = pw[i-1] * 2 % mod
	}
	for i := 0; i <= k; i++ {
		for j := 1; j <= m; j++ {
			if (i&(1<<(X[j]-1))) != 0 && (i&(1<<(Y[j]-1))) != 0 {
				s[i]++
			}
		}
	}
	for g := 1; g <= 2; g++ {
		for i := 0; i <= k; i++ {
			if (i & (1 << (g - 1))) == 0 {
				continue
			}
			f[g][i] = pw[s[i]]
			for j := i; j > 0; j = (j - 1) & i {
				f[g][i] = (f[g][i] - (f[g][i^j] * pw[s[j]] % mod) + mod) % mod
			}
		}
	}
	ans = pw[m]
	for i := 0; i <= k; i++ {
		for j := k ^ i; j > 0; j = (j - 1) & (k ^ i) {
			if s[i]+s[j] == s[i|j] {
				ans = (ans - ((f[1][i] * f[2][j] % mod) * pw[s[k^(i|j)]] % mod) + mod) % mod
			}
		}
	}
	fmt.Println(ans)
}
