package main

import "fmt"

func main() {
	var S int
	fmt.Scan(&S)
	T := S
	t1 := (S - 2) / 4
	t1 %= mod
	t2 := S/2 - 1
	t2 %= mod

	S %= mod
	sum := 0
	sum += divMod((S-1)*(S-2)%mod*(S-3)%mod*(S-4)%mod*(S-5)%mod, 120)
	sum %= mod
	sum += 6 * (((S-1)*t1%mod - 2*t1%mod*(t1+1)%mod + mod) % mod) % mod
	sum %= mod
	sum += 3 * (((1-S+mod)%mod*t2%mod-divMod(t2*(t2+1)%mod*(2*t2%mod+1)%mod, 3)+mod)%mod + divMod((S+1)%mod*t2%mod*(t2+1)%mod, 2))
	sum %= mod
	if T%3 == 0 {
		sum += 8 * (divMod(S, 3) - 1) % mod
		sum %= mod
	}
	if T%2 == 0 {
		sum += divMod(6*(divMod(S, 2)-1)%mod*(divMod(S, 2)-2)%mod, 2)
		sum %= mod
	}
	fmt.Println(divMod(sum, 24))
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
