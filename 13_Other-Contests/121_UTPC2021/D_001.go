package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353
const N = 100

var n int
var e [N + 5][]int
var C, g, s, t [N + 5][N + 5]int
var f [N + 5][N + 5][N + 5]int
var sz [N + 5]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}
	prep(n)
	dp(1, 0)
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ans += f[1][i][j]
		}
	}
	fmt.Println(ans % MOD)
}

func prep(n int) {
	C[0][0] = 1
	for i := 1; i <= n; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % MOD
		}
	}
}

func dp(u, tf int) {
	sz[u] = 1
	f[u][1][1] = 1
	for _, v := range e[u] {
		if v == tf {
			continue
		}
		dp(v, u)
		s1 := sz[u]
		s2 := sz[v]
		for i := 1; i <= s2; i++ {
			for j := 1; j <= s2; j++ {
				s[i][j] = (f[v][i][j] + s[i][j-1]) % MOD
			}
		}
		for i := 1; i <= s2; i++ {
			for j := 1; j <= s2; j++ {
				s[i][j] = (s[i][j] + s[i-1][j]) % MOD
			}
		}
		for i := 0; i <= s2; i++ {
			for j := 0; j <= s2; j++ {
				t[i][j] = s[i][s2] + s[s2][j] - 2*s[i][j]
			}
		}
		for i := 1; i <= s1; i++ {
			for j := 1; j <= s1; j++ {
				for k := 0; k <= s2; k++ {
					for l := 0; l <= s2; l++ {
						v := cal(i, k, s1-i, s2-k) * cal(j, l, s1-j, s2-l) % MOD
						g[i+k][j+l] = (g[i+k][j+l] + v*t[k][l]%MOD*f[u][i][j]) % MOD
					}
				}
			}
		}
		sz[u] += sz[v]
		for i := 1; i <= sz[u]; i++ {
			for j := 1; j <= sz[u]; j++ {
				if g[i][j] < 0 {
					g[i][j] += MOD
				}
				f[u][i][j] = g[i][j]
				g[i][j] = 0
			}
		}
	}
}

func cal(a, b, c, d int) int {
	return C[a+b-1][a-1] * C[c+d][c] % MOD
}
