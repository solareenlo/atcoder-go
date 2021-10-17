package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	mod := int(1e9 + 7)
	nCk := nCrMod(n*m-2, k-2, mod)
	res := 0
	for d := 1; d < n; d++ {
		res += d * (n - d) % mod * m % mod * m % mod
		res %= mod
	}
	for d := 1; d < m; d++ {
		res += d * (m - d) % mod * n % mod * n % mod
		res %= mod
	}
	res *= nCk
	fmt.Println(res % mod)
}

func powMod(a, n, mod int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a, mod int) int {
	return powMod(a, mod-2, mod)
}

func nCrMod(n, r, mod int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	res := int(1)
	for i := int(0); i < r; i++ {
		res = res * (n - i) % mod
		res = res * invMod(i+1, mod) % mod
	}
	return res
}
