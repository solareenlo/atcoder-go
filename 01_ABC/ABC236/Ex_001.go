package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var m int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n, &m)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	const mod = 998244353
	f := make([]int, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] * (1 - i) % mod
	}

	lc := make([]int, 1<<16)
	lc[0] = 1
	for s := 1; s < (1 << n); s++ {
		x := (s & -s)
		p := int(math.Log2(float64(x)))
		if lc[s^x] == -1 {
			lc[s] = -1
		} else {
			lc[s] = lcm(lc[s^x], a[p])
		}
	}

	p := make([]int, 1<<16)
	for s := 0; s < (1 << n); s++ {
		if lc[s] != -1 {
			x := 0
			for i := 0; i < n; i++ {
				x += ((s >> i) & 1)
			}
			p[s] = (m / lc[s]) % mod * f[x] % mod
		}
	}

	dp := make([]int, 1<<16)
	dp[0] = 1
	for s := 1; s < (1 << n); s++ {
		for j := s; j > 0; j = (j - 1) & s {
			if j&(s&-s) != 0 {
				dp[s] += dp[s^j] * p[j] % mod
				dp[s] %= mod
			}
		}
	}

	fmt.Println((dp[(1<<n)-1] + mod) % mod)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	a = a / gcd(a, b)
	if a > m/b {
		return -1
	}
	a *= b
	if a > m {
		return -1
	}
	return a
}
