package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 16
	d := make([]int, 1<<N)
	p := make([]int, 1<<N)
	d[0] = 1
	p[0] = 0

	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		b[i]++
		x := (mod - powMod(a[i], b[i])) % mod
		for j := 0; j < 1<<i; j++ {
			d[(1<<i)+j] = (d[j] * x) % mod
			p[(1<<i)+j] = p[j] + b[i]
		}
	}

	c := make([]int, N)
	for i := 0; i < n; i++ {
		x := invMod(a[i])
		c[i] = 1
		for j := 0; j < n; j++ {
			if j != i {
				y := (a[j] * x) % mod
				c[i] = c[i] * (mod + 1 - y)
				c[i] %= mod
			}
		}
		c[i] = invMod(c[i])
	}

	ans := 0
	for i := 0; i < 1<<n; i++ {
		if p[i] <= m {
			x := 0
			for j := 0; j < n; j++ {
				x = x + (powMod(a[j], m-p[i]) * c[j])
				x %= mod
			}
			ans = ans + (x * d[i])
			ans %= mod
		}
	}

	fmt.Println(ans)
}

const mod = 998244353

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
