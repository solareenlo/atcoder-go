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
	n *= 2

	const N = 610
	p := make([]int, N)
	for i := m; i > 0; i-- {
		var x, y int
		fmt.Fscan(in, &x, &y)
		p[x] = y
		p[p[x]] = x
	}

	const mod = 1_000_000_007
	g := make([]int, N)
	g[0] = 1
	for i := 2; i <= n; i += 2 {
		g[i] = g[i-2] * (i - 1) % mod
	}

	s := make([]int, N)
	for i := 1; i <= n; i++ {
		if p[i] == 0 {
			s[i] = s[i-1] + 1
		} else {
			s[i] = s[i-1]
		}
	}

	f := [N][N]int{}
	ans := 0
	for i := n; i > 0; i-- {
		for j := i + 1; j <= n; j += 2 {
			flg := 1
			for k := i; k <= j; k++ {
				if (p[k] >= i && p[k] <= j) || p[k] == 0 {
					flg &= 1
				} else {
					flg &= 0
				}
			}
			if flg == 0 {
				continue
			}
			f[i][j] = g[s[j]-s[i-1]]
			for k := i + 1; k < j; k += 2 {
				f[i][j] = (f[i][j]-f[i][k]*g[s[j]-s[k]])%mod + mod
				f[i][j] %= mod
			}
			ans += f[i][j] * g[n-m*2-s[j]+s[i-1]]
			ans %= mod
		}
	}

	fmt.Println(ans)
}
