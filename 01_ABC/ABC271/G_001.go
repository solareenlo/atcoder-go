package main

import (
	"bufio"
	"fmt"
	"os"
)

type qq struct {
	a [30][30]int
}

var zb, cs qq

func main() {
	in := bufio.NewReader(os.Stdin)

	var step, px, py int
	fmt.Fscan(in, &step, &px, &py)
	px, py = py, px

	var c string
	fmt.Fscan(in, &c)
	I := powMod(100, mod-2)
	px = px * I % mod
	py = py * I % mod

	for i := 0; i < 24; i++ {
		if c[i] == 'A' {
			cs.a[0][(i+1)%24] = 1
		}
	}

	ap := 1
	for i := 0; i < 24; i++ {
		tmp := 1 - py
		if c[i] == 'A' {
			tmp = 1 - px
		}
		ap = ap * tmp % mod
	}
	ap = powMod(1-ap, mod-2)

	for i := 0; i < 24; i++ {
		for j := 0; j < 24; j++ {
			p1 := py
			if c[j] == 'A' {
				p1 = px
			}
			tmp := j + 23
			if i <= j {
				tmp = j - 1
			}
			for k := i; k <= tmp; k++ {
				tmp1 := 1 - py
				if c[k%24] == 'A' {
					tmp1 = 1 - px
				}
				p1 = p1 * tmp1 % mod
			}
			zb.a[(j+1)%24][i] = p1 * ap % mod
		}
	}
	for ; step > 0; step >>= 1 {
		if (step & 1) != 0 {
			cs = mul(cs, zb)
		}
		zb = mul(zb, zb)
	}

	fmt.Println((cs.a[0][0] + mod) % mod)
}

func mul(a, b qq) qq {
	var c qq
	for k := 0; k < 24; k++ {
		for i := 0; i < 24; i++ {
			for j := 0; j < 24; j++ {
				c.a[i][j] += a.a[i][k] * b.a[k][j] % mod
				c.a[i][j] %= mod
			}
		}
	}
	return c
}

const mod = 998244353

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
