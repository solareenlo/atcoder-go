package main

import "fmt"

func main() {
	const N = 105

	var f [N][N][2]int

	var n, a, b, p, q int
	fmt.Scan(&n, &a, &b, &p, &q)
	f[a][b][0] = 1

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			for x := i + 1; x <= i+p; x++ {
				f[min(x, n)][j][1] = (f[min(x, n)][j][1] + divMod(f[i][j][0], p)) % MOD
			}
			for y := j + 1; y <= j+q; y++ {
				f[i][min(y, n)][0] = (f[i][min(y, n)][0] + divMod(f[i][j][1], q)) % MOD
			}
		}
	}
	s1, s2 := 0, 0
	for i := 1; i < n; i++ {
		s1 = (s1 + f[n][i][1]) % MOD
		s2 = (s2 + f[i][n][0]) % MOD
	}
	fmt.Println(divMod(s1, (s2+s1)%MOD))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const MOD = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a int) int {
	b, u, v := MOD, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
