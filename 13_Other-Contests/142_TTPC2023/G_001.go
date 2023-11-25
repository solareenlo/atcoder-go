package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	var C, inv [10000001]int

	C[0] = 1
	inv[1] = 1
	for i := 2; i <= N; i++ {
		inv[i] = (998244353 / i) * (MOD - inv[998244353%i]) % MOD
	}
	for d := 1; d < M; d++ {
		C[d] = (C[d-1] * (N + d) % MOD) * inv[d] % MOD
	}

	ans := 0
	for n := 0; ; n++ {
		d := n * (3*n - 1) / 2
		if d >= M {
			break
		}
		var coef int
		if n%2 != 0 {
			coef = MOD - 1
		} else {
			coef = 1
		}
		ans = (ans + (coef * C[M-d-1] % MOD)) % MOD
	}
	for n := -1; ; n-- {
		d := n * (3*n - 1) / 2
		if d >= M {
			break
		}
		var coef int
		if n%2 != 0 {
			coef = MOD - 1
		} else {
			coef = 1
		}
		ans = (ans + (coef * C[M-d-1] % MOD)) % MOD
	}
	for i := 1; i <= N; i++ {
		ans = ans * inv[i] % MOD
	}
	fmt.Println(ans)
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
