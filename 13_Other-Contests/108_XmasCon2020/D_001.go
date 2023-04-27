package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m, v int
var s [1000005]int
var f [1005]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var c int
	fmt.Fscan(in, &n, &c)

	if c == 0 || n == 1 {
		fmt.Println(1)
		return
	}
	if c == 1 {
		if n == 2 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		return
	}

	m = min(n, 1000000)
	v = c * powMod(c-1, mod-2) % mod
	s[1] = 1
	for i := 1; i <= m; i++ {
		if i > 1 {
			s[i] = s[i] * v % mod
		}
		for j := 2 * i; j <= m; j += i {
			s[j] = (s[j] - s[i] + mod) % mod
		}
		s[i] = (s[i] + s[i-1]) % mod
	}

	for i := range f {
		f[i] = -1
	}
	fmt.Println(solve(n) * powMod(mod-c+1, n-1) % mod)
}

func solve(q int) int {
	if q <= m {
		return s[q]
	}
	if f[n/q] != -1 {
		return f[n/q]
	}
	s := 0
	for l, r := 2, 0; l <= q; l = r + 1 {
		r = q / (q / l)
		s = (s + solve(q/l)*(r-l+1)) % mod
	}
	f[n/q] = (1 - s*v%mod + mod) % mod
	return f[n/q]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
