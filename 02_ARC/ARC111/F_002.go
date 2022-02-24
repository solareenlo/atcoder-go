package main

import "fmt"

func main() {
	var N, M, Q int
	fmt.Scan(&N, &M, &Q)

	C := divMod((2*M+1)%mod*N%mod*(N+1)%mod, 2)
	CQ := powMod(C, Q-1)
	ans := CQ * Q % mod * divMod(M*(M-1)%mod, 2) % mod * divMod(N*(N+1)%mod*(N+2)%mod, 6) % mod
	ret := CQ * C % mod * N % mod
	for i := 0; i < N; i++ {
		tmp := C - (M * (i + 1) % mod * (N - i) % mod)
		tmp += mod
		tmp %= mod
		ret -= powMod(tmp, Q)
		ret += mod
		ret %= mod
	}
	tmp := divMod(ret*(M-1)%mod, 2)
	ans = divMod(((ans-tmp)+mod)%mod, M)
	fmt.Println(ans)
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
