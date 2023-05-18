package main

import "fmt"

func main() {
	initMod()
	var h, w int
	fmt.Scan(&h, &w)
	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}

	ans := 0
	for i := 0; i < w; i++ {
		ans = (ans + nCrMod(h-2+i, i)*powMod(2, (h-1)*(w-1)-i)%MOD) % MOD
	}
	fmt.Println(ans)
}

const MOD = 998244353
const size = 400010

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
