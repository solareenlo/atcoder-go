package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	initMod()

	s := make([]int, N)
	pw := make([]int, N)
	s[0] = 1
	pw[0] = 1
	for i := 1; i < N; i++ {
		pw[i] = pw[i-1] * 3 % mod
	}

	ans := pw[m+k]
	for i := 1; i <= m+k; i++ {
		s[i] = 2 * s[i-1] % mod
		if i > m {
			s[i] += mod - nCrMod(i-1, m)
			s[i] %= mod
		}
		if i > k {
			s[i] += mod - nCrMod(i-1, i-k-1)
			s[i] %= mod
		}
		ans += nCrMod(n+i-1, i) * s[i] % mod * pw[m+k-i]
		ans %= mod
	}
	fmt.Println(ans)
}

const mod = 1000000007
const N = 1010101

var fact, invf [N]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < N; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
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

func invMod(a int) int {
	return powMod(a, mod-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
