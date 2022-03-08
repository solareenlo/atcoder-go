package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353
const two = 499122177
const six = 166374059

func s1(l, r int) int {
	if l > r {
		return 0
	}
	n := (r - l + 1) % mod
	return (l + r) % mod * n % mod * two % mod
}

func s2(n int) int {
	n %= mod
	return n * (n + 1) % mod * (n + n + 1) % mod * six % mod
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)

	for l := 0; l < T; l++ {
		ans := 0
		var n, x int
		fmt.Fscan(in, &n, &x)
		for i := 1; i < x && i <= n; i++ {
			k := (x - 1) / (i + 1)
			if k == 0 {
				ans += x % mod * (i % mod)
				ans %= mod
			} else {
				t := min((x-1-k*(i+1))/(2*k), n-i)
				ans += (t+1)%mod*(x%mod)%mod*(i%mod) + (x-k*i)%mod*s1(1, t) - k%mod*s2(t)%mod + mod
				ans %= mod
				i += t
				x -= (t + 1) * k
			}
		}
		ans += x % mod * s1(x, n)
		ans %= mod
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
