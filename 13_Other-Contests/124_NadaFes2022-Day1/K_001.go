package main

import "fmt"

func main() {
	var N, L, I int
	fmt.Scan(&N, &L, &I)
	if L%2 == 1 {
		fmt.Println(0)
		return
	}

	initMod()

	ans := 0

	for i := I % 2; i <= I; i += 2 {
		ans = (ans + nCrMod((I-i)/2+L/2, L/2)*nCrMod(L+(I+i)/2, i)%MOD) % MOD
	}

	fmt.Println(ans * powMod(2, L/2) % MOD)
}

const MOD = 998244353
const size = 20000001

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
