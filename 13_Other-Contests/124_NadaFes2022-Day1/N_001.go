package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	var N, M int
	fmt.Fscan(in, &N, &M)
	B := make([]int, N)
	for i := range B {
		fmt.Fscan(in, &B[i])
	}
	B = append(B, 0)
	even := 0
	odd := 0
	last := 0
	for _, b := range B {
		if last != b {
			odd++
		} else {
			even++
		}
		last = b
	}
	ret := 0
	n := N + 1
	k := M
	d := odd
	for p := 0; p <= n; p++ {
		tmp1 := 0
		tmp2 := 0
		for i := 0; i < d+1; i++ {
			if (d-i)%2 != 0 {
				tmp1 = (tmp1 + (nCrMod(p, i) * ((MOD - nCrMod(n-p, d-i)) % MOD) % MOD)) % MOD
			} else {
				tmp1 = (tmp1 + (nCrMod(p, i) * nCrMod(n-p, d-i) % MOD)) % MOD
			}
		}
		for i := 0; i < k+1; i++ {
			if i%2 != 0 {
				tmp2 = (tmp2 + nCrMod(p*(n-p), i)*((MOD-nCrMod(n*(n-1)/2-p*(n-p), k-i))%MOD)%MOD) % MOD
			} else {
				tmp2 = (tmp2 + (nCrMod(p*(n-p), i) * nCrMod(n*(n-1)/2-p*(n-p), k-i) % MOD)) % MOD
			}
		}
		ret = (ret + ((nCrMod(n, p) * tmp1 % MOD) * tmp2 % MOD)) % MOD
	}
	fmt.Println(divMod(ret, (powMod(2, n)*nCrMod(n, odd))%MOD))
}

const MOD = 998244353
const size = 5000000

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, MOD-2)
}

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % MOD * invf[n-r] % MOD
}

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
