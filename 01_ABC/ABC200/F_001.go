package main

import "fmt"

func main() {
	var S string
	var K int
	fmt.Scan(&S, &K)

	N := len(S)
	q := 0
	for i := range S {
		if S[i] == '?' {
			q++
		}
	}

	inv2 := modInv(2)
	res := 0
	for i := 0; i < N; i++ {
		if S[i]+S[(i+1)%N] > 'b' {
			res += inv2
			res %= mod
		} else if S[i] != S[(i+1)%N] {
			res++
		}
	}

	res *= inv2 * K % mod * powMod(2, K*q) % mod
	res %= mod
	if K*N > 1 {
		fmt.Println(res)
	} else {
		fmt.Println(0)
	}
}

const mod = 1000000007

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
