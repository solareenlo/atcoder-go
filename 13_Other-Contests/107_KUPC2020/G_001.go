package main

import "fmt"

func main() {
	const N = 3007

	var n, m int
	fmt.Scan(&n, &m)
	n++
	m--

	var f, c, inv [N]int
	c[0] = 1
	fac := 1
	for i := 1; i <= n; i++ {
		inv[i] = powMod(i, mod-2)
		c[i] = c[i-1] * (m - i) % mod * inv[i] % mod
		fac = fac * i % mod
	}
	c[0] = 2
	f[1] = inv[n]
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			f[i] = (f[i] + f[j]*c[i-j-1]%mod*inv[n-j]) % mod
		}
	}
	fmt.Println(f[n] * fac % mod)
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
