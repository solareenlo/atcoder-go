package main

import (
	"bufio"
	"fmt"
	"os"
)

const mx = 100000
const N = 200005
const MOD = 998244353

var s, mul, mu, p, b, Inv [N]int
var cnt int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < mx+1; i++ {
		mul[i] = 1
	}
	for i := 1; i < mx+1; i++ {
		Inv[i] = kuai(i, MOD-2)
	}
	for i := 1; i < n+1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		solve(1, mx, 0, a, 0, b, a, b)
	}
	work(mx)
	for i := 1; i < mx+1; i++ {
		mul[i] = mul[i-1] * mul[i] % MOD
		s[i] += s[i-1]
	}
	for i := 1; i < mx+1; i++ {
		if s[i] != 0 {
			mul[i] = 0
		}
	}
	for i := 1; i < m+1; i++ {
		s := 0
		for j := 1; j <= mx/i; j++ {
			s = (s + mul[i*j]*mu[j]%MOD) % MOD
		}
		fmt.Fprintln(out, s)
	}
}

func kuai(a, b int) int {
	ans := 1
	for ; b > 0; b, a = b>>1, a*a%MOD {
		if (b & 1) != 0 {
			ans = ans * a % MOD
		}
	}
	return ans
}

func solve(l, r, la, ra, lb, rb, A, B int) {
	if l > r {
		return
	}
	if l == r {
		t := B/l - A/l
		if t == 0 {
			s[l]++
			s[r+1]--
		} else {
			mul[l] = mul[l] * t % MOD
			mul[r+1] = mul[r+1] * Inv[t] % MOD
		}
		return
	}
	if la == ra && lb == rb {
		t := lb - la
		if t == 0 {
			s[l]++
			s[r+1]--
		} else {
			mul[l] = mul[l] * t % MOD
			mul[r+1] = mul[r+1] * Inv[t] % MOD
		}
		return
	}
	mid := (l + r) >> 1
	solve(l, mid, A/mid, ra, B/mid, rb, A, B)
	solve(mid+1, r, la, A/mid, lb, B/mid, A, B)
}

func work(maxn int) {
	mu[1] = 1
	for i := 2; i <= maxn; i++ {
		if b[i] == 0 {
			cnt++
			p[cnt] = i
			mu[i] = MOD - 1
		}
		for j := 1; j <= cnt && p[j]*i <= maxn; j++ {
			b[i*p[j]] = 1
			if i%p[j] == 0 {
				break
			}
			mu[i*p[j]] = MOD - mu[i]
		}
	}
}
