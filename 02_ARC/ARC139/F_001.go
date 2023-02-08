package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	N := max(n, m) + 1

	initMod(N)

	nw := 1
	res := 0
	for i := 1; i <= min(n, m); i++ {
		nw = nw * ((_k[n] - _k[i-1] + mod) % mod) % mod
		tmp1 := (_k[m+1] - 1 - _k[m-i] + mod)
		tmp2 := (i*_k[i] + 1) % mod
		res = (res + (tmp1*nCrMod(m, i)%mod-tmp2*nCrMod(m, i+1)%mod+mod)%mod*nw) % mod
	}
	fmt.Println(res * inv2 % mod)
}

const inv2 = (mod + 1) / 2

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const mod = 998244353

var fact, invf, _k []int

func initMod(size int) {
	fact = make([]int, size+1)
	invf = make([]int, size+1)
	_k = make([]int, size+1)
	_k[0] = 1
	for i := 1; i <= size; i++ {
		_k[i] = (_k[i-1] + _k[i-1]) % mod
	}
	fact[0] = 1
	for i := 1; i <= size; i++ {
		fact[i] = (fact[i-1] * (_k[i] - 1)) % mod
	}
	invf[0] = 1
	invf[size] = invMod(fact[size])
	for i := size; i > 0; i-- {
		invf[i-1] = invf[i] * (_k[i] - 1) % mod
	}
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

func invMod(a int) int {
	return powMod(a, mod-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
