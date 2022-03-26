package main

import "fmt"

func main() {
	initMod()

	var N, M int
	fmt.Scan(&N, &M)

	if N > M {
		N ^= M
		M ^= N
		N ^= M
	}

	ans := M
	tot := invf[N+M] * fact[N] % mod * fact[M] % mod
	for i := 1; i <= N; i++ {
		ans = ans + tot*nCrMod(N-i, M-i)%mod*nCrMod(i, i)%mod*I[2]
		ans %= mod
	}
	fmt.Println(ans)
}

const mod = 119<<23 | 1
const size = 2000000

var fact, invf, I [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
	I[1] = 1
	for i := 2; i < size; i++ {
		I[i] = (mod - mod/i) * I[mod%i] % mod
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
	return fact[n+r] * invf[n] % mod * invf[r] % mod
}
