package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N = 60
const mod = 998244353

var (
	n  int
	m  int
	f  = [N * N]int{}
	g  = [N * N]int{}
	h  = [N * N]int{}
	a  = [N]int{}
	dp = [N][N]int{}
	s  = make([][]string, N)
)

func sol() int {
	t := 0
	for i := 0; i <= n*m; i++ {
		f[i] = 0
	}
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= m; k++ {
				dp[j][k] = 0
			}
		}
		dp[0][0] = 1
		for j := 1; j <= m; j++ {
			if s[i][j] == "0" {
				for a := j - 1; a >= 0; a-- {
					for b := j - 1; b >= 0; b-- {
						dp[a+1][b] += dp[a][b]
						dp[a+1][b] %= mod
						dp[a+1][b+1] -= dp[a][b]
						dp[a+1][b+1] += mod
						dp[a+1][b+1] %= mod
						dp[a][b+1] += dp[a][b]
						dp[a][b+1] %= mod
					}
				}
			} else {
				for a := j - 1; a >= 0; a-- {
					for b := j - 1; b >= 0; b-- {
						dp[a+1][b] += dp[a][b]
						dp[a+1][b] %= mod
						dp[a+1][b+1] += dp[a][b]
						dp[a+1][b+1] %= mod
						dp[a][b+1] -= dp[a][b]
						dp[a][b+1] += mod
						dp[a][b+1] %= mod
					}
				}
			}
		}
		for j := 0; j <= m; j++ {
			g[j] = dp[a[i]][j]
		}
		for i := 0; i <= t; i++ {
			for j := 0; j <= m; j++ {
				h[i+j] += f[i] * g[j] % mod
				h[i+j] %= mod
			}
		}
		t += m
		for i := 0; i <= t; i++ {
			f[i] = h[i]
			h[i] = 0
		}
	}
	z := 0
	for i := 1; i <= t; i++ {
		z += divMod(f[i], i)
		z %= mod
	}
	z = divMod(z*n%mod*m%mod, 2)
	return z
}

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

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		var S string
		fmt.Fscan(in, &S)
		S = " " + S
		s[i] = strings.Split(S, "")
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	x := sol()
	k := f[0]
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j <= a[i] {
				s[i][j] = "1"
			} else {
				s[i][j] = "0"
			}
		}
	}
	y := sol()
	ans := divMod((y-x+mod)%mod, k)
	fmt.Println(ans)
}
