package main

import "fmt"

func main() {
	initMod()

	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	if C == 0 || D == 0 {
		fmt.Println(nCrMod(A+B, A))
		return
	}

	ans := 0
	for a := C; a <= A; a++ {
		for b := D; b <= B; b++ {
			t := nCrMod(a-C+D-1, D-1) * nCrMod(b-D+C-1, C-1) % mod * nCrMod(C+D, C) % mod * nCrMod(A+B-a-b, A-a) % mod
			ans = (ans + t) % mod
		}
	}
	fmt.Println(ans)
}

const mod = 1000000007
const size = 10000

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
