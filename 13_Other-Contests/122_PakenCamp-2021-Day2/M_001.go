package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	now := 0
	way := 1
	for i := 0; i < N; i++ {
		sum := 0
		if i%2 == 0 {
			n := int(mint(i).div(2))
			sum = (n*(n+1)%MOD - n + MOD) % MOD
		} else {
			n := int(mint(i - 1).div(2))
			sum = n * (n + 1) % MOD
		}
		now = (now*(i+1)%MOD + way*sum%MOD) % MOD
		way = way * (i + 1) % MOD
	}
	fmt.Println(now)
}

type mint int

func (m mint) pow(p int) mint {
	return powMod(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return divMod(m, n)
}

const MOD = 998244353

func powMod(a mint, n int) mint {
	res := mint(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a mint) mint {
	return powMod(a, MOD-2)
}

func divMod(a, b mint) mint {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a mint) mint {
	b, u, v := mint(MOD), mint(1), mint(0)
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
