package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var c [255][255]int
	for i := 0; i < 255; i++ {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % MOD
		}
	}
	var g, f, h [35][255]int
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= j*(i-j); k++ {
				g[i][k] = (g[i][k] + c[i][j]*c[j*(i-j)][k]) % MOD
			}
		}
	}
	f[1][0] = 2
	for i := 2; i <= n; i++ {
		for j := i - 1; j <= i*i/4; j++ {
			f[i][j] = g[i][j]
			for k := 1; k <= i; k++ {
				for l := k - 1; l <= min(k*k/4, j); l++ {
					f[i][j] = (f[i][j] - c[i-1][k-1]*f[k][l]%MOD*g[i-k][j-l]%MOD + MOD) % MOD
				}
			}
		}
	}
	h[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i*i/4; j++ {
			for k := 1; k <= i; k++ {
				for l := k - 1; l <= min(k*k/4, j); l++ {
					h[i][j] = (h[i][j] + c[i-1][k-1]*h[i-k][j-l]%MOD*f[k][l]%MOD*(MOD+1)/2) % MOD
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= n*n/4; i++ {
		for j, x := 0, 1; j < i; j, x = j+1, -x {
			ans = (ans + h[n][i]*c[i][j]%MOD*powMOD(i-j, m)%MOD*x + MOD) % MOD
		}
	}
	fmt.Println(ans * powMOD(2, m) % MOD)
}

const MOD = 998244353

func powMOD(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
