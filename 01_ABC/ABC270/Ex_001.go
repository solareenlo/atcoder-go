package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	y := make([]int, n)
	s := 0
	for i := n - 1; i >= 1; i-- {
		x := powMod(divMod(n, i), a[i]-a[i-1])
		z := divMod((n-s+mod)%mod, (n-i+mod)%mod)
		y[i-1] = (((y[i]+z)%mod)*x%mod - z + mod) % mod
		s += y[i-1]
		s %= mod
	}
	fmt.Println(y[0])
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
