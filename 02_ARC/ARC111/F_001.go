package main

import "fmt"

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	res := 0
	t := powMod(n*(n+1)*(2*m+1)%mod, mod-2)
	for i := 1; i <= n; i++ {
		pi := 2 * i * (n - i + 1) % mod * t % mod * m % mod
		s := (q - (1-powMod(1-pi+mod, q)+mod)*powMod(pi, mod-2)%mod + mod) % mod
		res = (res + t*(m-1)%mod*i%mod*(n-i+1)%mod*s) % mod
	}

	iv2 := (mod + 1) / 2
	fmt.Println(res * powMod(n*(n+1)%mod*iv2%mod*(2*m+1)%mod, q) % mod)
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
