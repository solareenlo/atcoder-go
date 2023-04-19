package main

import "fmt"

func main() {
	var x, y, N int
	fmt.Scan(&x, &y, &N)

	initMod()

	ans := 0
	for k := 1; k <= N-1; k++ {
		if 2*k >= N {
			ans = (ans + x*nCrMod(k-1, N-1-k)%mod*invf[k]%mod) % mod
		}
		if 2*k >= N-1 {
			ans = (ans + y*nCrMod(k, N-1-k)%mod*invf[k]%mod) % mod
		}
	}
	fmt.Println(ans * fact[N-1] % mod)
}

const mod = 1000000007
const size = 1000100

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
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
