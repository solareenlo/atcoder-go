package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1005
const mod = 998244353

var (
	n    int
	f    = [N]int{}
	g    = [N]int{}
	p    = [N]int{}
	s    = [N]int{}
	fact = [N]int{}
	inv  = [N]int{}
	m    int
	a    = [N]int{}
	l    = [N]int{}
)

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

func cz(x int) int {
	p[0] = 1
	for i := 1; i <= n+2; i++ {
		p[i] = p[i-1] * (x - i + mod) % mod
	}
	s[n+3] = 1
	for i := n + 2; i >= 1; i-- {
		s[i] = s[i+1] * (x - i + mod) % mod
	}
	r := 0
	for i := 1; i <= n+2; i++ {
		v := p[i-1] * s[i+1] % mod * f[i] % mod * inv[i-1] % mod * inv[n+2-i] % mod
		if (n+i)&1 != 0 {
			v = (mod - v) % mod
		}
		r = (r + v) % mod
	}
	return r
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l[n+1] = m
	fact[0] = 1
	inv[0] = 1
	for i := 1; i <= n+2; i++ {
		fact[i] = fact[i-1] * i % mod
		inv[i] = powMod(fact[i], mod-2)
	}

	for i := 1; i <= n+2; i++ {
		f[i] = 1
	}

	for i := n; i >= 1; i-- {
		l[i] = l[i+1] / a[i]
		for j := 1; j <= n+2; j++ {
			f[j] = (f[j] + f[j-1]) % mod
		}
		r := cz(l[i+1] % mod)
		if a[i] > 1 {
			for j := 1; j <= n+2; j++ {
				g[j] = (r + mod - cz((a[i]*j-1)%mod)) % mod
			}
		} else {
			for j := 1; j <= n+2; j++ {
				g[j] = (r + mod - f[j-1]) % mod
			}
		}
		for j := 1; j <= n+2; j++ {
			f[j] = g[j]
		}
	}

	for j := 1; j <= n+2; j++ {
		f[j] = (f[j] + f[j-1]) % mod
	}
	fmt.Println(cz(l[1] % mod))
}
