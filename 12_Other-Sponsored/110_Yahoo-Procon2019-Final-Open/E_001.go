package main

import "fmt"

type mat struct {
	n, m int
	a    [201][201]int
}

func (a *mat) mul(b mat) mat {
	var c mat
	c.n = a.n
	c.m = b.m
	for i := 0; i < a.n+1; i++ {
		for k := 0; k < a.m+1; k++ {
			if a.a[i][k] != 0 {
				for j := 0; j <= b.m; j++ {
					c.a[i][j] = (c.a[i][j] + a.a[i][k]*b.a[k][j]) % MOD
				}
			}
		}
	}
	return c
}

func main() {
	var fac, ifac [401]int
	var c, g [401][401]int
	var n, m int
	fmt.Scan(&n, &m)
	m--
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i % MOD
	}
	ifac[n] = powMod(fac[n], MOD-2)
	for i := n - 1; i >= 0; i-- {
		ifac[i] = ifac[i+1] * (i + 1) % MOD
	}
	c[0][0] = 1
	for i := 1; i <= n; i++ {
		c[i][0] = 1
		c[i][i] = 1
		for j := 1; j < i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % MOD
		}
	}
	for i := 0; i < n+1; i++ {
		g[i][0] = 1
		for j := 1; j+i <= n; j++ {
			g[i][j] = g[i][j-1] * (i + j - 1) % MOD
		}
	}
	var a, b mat
	a.n = 1
	a.m = n / 2
	k := a.m
	b.m = k
	b.n = b.m
	for i := 0; i < k+1; i++ {
		for j := 1; j <= k; j++ {
			for k := 0; k < j+1; k++ {
				if (k & 1) != 0 {
					b.a[i][j] = (b.a[i][j] - (c[n-2*i][k]*c[n-i-k][j-k]%MOD)*g[j-k][n-i-j]%MOD + MOD) % MOD
				} else {
					b.a[i][j] = (b.a[i][j] + c[n-2*i][k]*c[n-i-k][j-k]%MOD*g[j-k][n-i-j]) % MOD
				}
			}
		}
	}
	a.a[1][0] = 1
	for m > 0 {
		if (m & 1) != 0 {
			a = a.mul(b)
		}
		b = b.mul(b)
		m >>= 1
	}
	ans := 0
	for i := 1; i <= k; i++ {
		ans = (ans + a.a[1][i]*fac[n-i-1]) % MOD
	}
	fmt.Println(ans)
}

const MOD = 1000000007

func powMod(a, n int) int {
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
