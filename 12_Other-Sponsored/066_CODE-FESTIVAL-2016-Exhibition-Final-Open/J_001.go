package main

import "fmt"

func main() {
	initMod()

	var n, A, B, C int
	fmt.Scan(&n, &A, &B, &C)
	if B&1 != 0 {
		fmt.Println(0)
		return
	}
	B /= 2
	ans := 0
	for x := 0; x <= C; x++ {
		for p := 0; p <= (C-x)/3; p++ {
			ans = (ans + ((nCrMod(A, C-x-3*p)*nCrMod(A+p, p)%MOD)*nCrMod(A+p+B, A+p)%MOD)*nCrMod(B+x-1, B-1)%MOD) % MOD
			ans %= MOD
		}
	}
	fmt.Println(ans)
}

const MOD = 1000000007
const size = 101010

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
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if n < 0 || r < 0 {
		return 0
	}
	return (fact[n] * invf[r] % MOD) * invf[n-r] % MOD
}
