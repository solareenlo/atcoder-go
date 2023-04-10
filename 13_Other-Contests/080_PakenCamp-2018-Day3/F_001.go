package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--

		var a, b int
		fmt.Fscan(in, &a, &b)
		if a == b {
			fmt.Fprintln(out, 1)
			continue
		}

		var getA func(int) int
		getA = func(x int) int { return x*a + x*(x-1)/2 }
		var getB func(int) int
		getB = func(x int) int { return x*b - x*(x-1)/2 }

		l := 0
		r := (b - a + 2) / 2
		for l+1 < r {
			m := (l + r) >> 1
			if getA(m+1) <= getB(m) {
				r = m
			} else {
				l = m
			}
		}

		res := 0
		res = (res - divMod(((l)*(l+1)%mod)*(2*l+1)%mod, (6)) + mod) % mod
		res = (res + divMod(((b-a+1)*(l)%mod)*(l+1)%mod, (2))) % mod
		res = (res + (l + 1)) % mod
		res = (res * (2)) % mod
		if r < (b-a+2)/2 || (b-a+1)%2 == 0 {
			res = (res + (getB(b-a+1-r)-getA(r)+1+mod)%mod) % mod
		}
		res = (res - (1) + mod) % mod
		fmt.Fprintln(out, res)
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
