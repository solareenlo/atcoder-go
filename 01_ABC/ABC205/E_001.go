package main

import "fmt"

func path(X, Y int) int {
	if X < 0 || Y < 0 {
		return 0
	}
	ret := 1
	for i := 1; i < X+Y+1; i++ {
		ret *= i
		ret %= mod
	}
	for i := 1; i < X+1; i++ {
		ret = divMod(ret, i)
	}
	for i := 1; i < Y+1; i++ {
		ret = divMod(ret, i)
	}
	return ret
}

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)

	if N > M+K {
		fmt.Println(0)
	} else {
		fmt.Println((path(N, M) - path(N-K-1, M+K+1) + mod) % mod)
	}
}

const mod = 1000000007

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
