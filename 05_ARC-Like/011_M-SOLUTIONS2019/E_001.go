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

	f := make([]int, mod)
	f[0] = 1
	for i := 1; i < mod; i++ {
		f[i] = f[i-1] * i % mod
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x, d, n int
		fmt.Fscan(in, &x, &d, &n)
		if d != 0 {
			t := divMod(x, d)
			if t+n > mod {
				fmt.Fprintln(out, 0)
			} else {
				fmt.Fprintln(out, divMod(f[t+n-1], f[t-1])*powMod(d, n)%mod)
			}
		} else {
			fmt.Fprintln(out, powMod(x, n))
		}
	}
}

const mod = 1000003

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
