package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	d := k + m
	b := make([]int, d+1)
	s := make([]int, d+2)
	s[d+1] = 1

	for i := 0; i < d+1; i++ {
		b[i] = powMod(i, k)
	}
	for j := 0; j < m; j++ {
		for i := 0; i < d; i++ {
			b[i+1] += b[i]
			b[i+1] %= mod
		}
	}

	x := n % mod
	for i := d; i+1 > 0; i-- {
		s[i] = s[i+1] * (x - i + mod) % mod
	}

	t := 1
	for i := 1; i < d+1; i++ {
		t *= (-i + mod) % mod
		t %= mod
	}
	t = invMod(t)

	y := 0
	for i := 1; i < d+2; i++ {
		y += t * s[i] % mod * b[i-1] % mod
		y %= mod
		t *= divMod(((x-i+1)*(-(d-i+1)+mod)%mod)%mod, i)
		t %= mod
	}
	fmt.Println(y)
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

func invMod(a int) int {
	return powMod(a, mod-2)
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
