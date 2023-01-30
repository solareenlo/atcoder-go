package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1000005

	var n, m int
	fmt.Fscan(in, &n, &m)
	var a, cnt [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		cnt[a[i]] += a[i]
	}
	var f [N]int
	f[0] = 1
	var g [11][N]int
	for i := 1; i <= 10; i++ {
		g[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		res := 0
		for j := 1; j <= min(10, i); j++ {
			if cnt[j] != 0 {
				res = (res + cnt[j]*g[j][i-j]) % mod
			}
		}
		f[i] = res * powMod(i, mod-2) % mod
		for j := 1; j <= 10; j++ {
			g[j][i] = f[i]
		}
		for j := 1; j <= min(10, i); j++ {
			g[j][i] = (g[j][i] - g[j][i-j] + mod) % mod
		}
	}
	lst := f[m]
	f[0] = 1
	for i := 1; i <= 10; i++ {
		g[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		res := 0
		for j := 1; j <= min(10, i); j++ {
			if cnt[j] != 0 {
				res = (res + cnt[j]*g[j][i-j]) % mod
			}
		}
		f[i] = mod - res*powMod(i, mod-2)%mod
		for j := 1; j <= 10; j++ {
			g[j][i] = f[i]
		}
		for j := 1; j <= min(10, i); j++ {
			g[j][i] = (g[j][i] + g[j][i-j]) % mod
		}
	}
	fmt.Println((lst - f[m] + mod) * powMod(2, mod-2) % mod)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
