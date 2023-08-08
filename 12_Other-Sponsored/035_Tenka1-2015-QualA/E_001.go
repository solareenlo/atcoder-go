package main

import "fmt"

func main() {
	const N = 409
	const MOD = 1000000007

	var n, m int
	fmt.Scan(&n, &m)
	var f [N]int
	var c, g [N][N]int
	f[0] = 1
	g[0][0] = 1
	for i := 0; i < n+1; i++ {
		c[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % MOD
		}
	}
	for i := m; i <= n; i++ {
		for j := 0; j <= i-m; j++ {
			f[i] = (f[i] + f[j]*c[i-1][j]%MOD) % MOD
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			for k := 0; k < i+1; k++ {
				g[i+1][j+1] = (g[i+1][j+1] + g[i-k][j]*f[k]%MOD) % MOD
			}
		}
	}
	for i := m; i <= n; i++ {
		for j := m; j < i; j++ {
			f[i] = (f[i] - g[i][j]*f[j]%MOD + MOD) % MOD
		}
	}
	fmt.Println(f[n])
}
