package main

import (
	"fmt"
)

const mod int = 998244353

type P struct {
	first  int
	second int
}

func fpow(a int, b int) int {
	ret := 1
	for ; b > 0; b >>= 1 {
		if (b & 1) != 0 {
			ret = ret * a % mod
		}
		a = a * a % mod
	}
	return ret
}

func calc(x int) P {
	if x == 0 {
		return P{0, 1}
	}
	xx := calc(x >> 1)
	a := xx.first
	b := xx.second
	c := ((2*b - a + mod) % mod) * a % mod
	d := (a*a%mod + b*b%mod) % mod
	if (x & 1) != 0 {
		return P{d, (c + d) % mod}
	}
	return P{c, d}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if m == 1 {
		fmt.Println(1)
		return
	}
	ans := calc(m + 2*(n-1)).first
	c := 1
	for i := 1; i < n; i++ {
		ans = (ans + mod - c*calc(2*n-2*i).first%mod) % mod
		c = (((m + i - 1) % mod) * c % mod) * fpow(i, mod-2) % mod
	}
	fmt.Println(ans)
}
