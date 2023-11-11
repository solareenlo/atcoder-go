package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353
const N = 2005

var f, g, h [2][N]int
var s [2]string

func add(x, y, q, p int) {
	f[x][y] = (f[x][y] + f[q][p]) % MOD
	g[x][y] = (g[x][y] + g[q][p]) % MOD
	h[x][y] = (h[x][y] + h[q][p]) % MOD
	if len(s[x]) > 0 && s[x][y] == 'Y' && len(s[q]) > 0 && s[q][p] == 'Y' {
		f[x][y] = (f[x][y] + (2*g[q][p]%MOD+h[q][p])%MOD) % MOD
		g[x][y] = (g[x][y] + h[q][p]) % MOD
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	h[1][1] = 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i&1])
		s[i&1] = " " + s[i&1] + " "
		for j := 1; j <= m; j++ {
			if i == 1 && j == 1 {
				continue
			}
			k := i & 1
			l := k ^ 1
			f[k][j] = 0
			g[k][j] = 0
			h[k][j] = 0
			add(k, j, l, j)
			add(k, j, k, j-1)
		}
	}
	fmt.Println(f[n&1][m])
}
