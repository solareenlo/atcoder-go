package main

import "fmt"

func main() {
	const N = 2020
	var C [N][N]int
	for i := 0; i < N; i++ {
		for j := 0; j < i+1; j++ {
			if j != 0 {
				C[i][j] = (C[i-1][j] + C[i-1][j-1]) % mod
			} else {
				C[i][j] = 1
			}
		}
	}
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	var c [N]int
	for i := 1; i <= n; i++ {
		var x int
		fmt.Scan(&x)
		c[x]++
	}
	ans := 0
	for i, w, t := 0, c[0], 0; i < m; t, i = t+c[i+1], i+1 {
		for j := 0; j <= min(k-t-1, w); j++ {
			ans = (ans + ((C[w][j] * powMod(i, j) % mod) * powMod(m-i, w-j) % mod)) % mod
		}
	}
	fmt.Println(divMod(ans, powMod(m, c[0])))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

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
