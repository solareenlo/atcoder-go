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

	var n, m int
	fmt.Fscan(in, &n, &m)

	var u, v [510]int
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &u[i], &v[i])
		u[i]--
		v[i]--
	}

	f := make([]int, 1<<14)
	fac := make([]int, 14)
	f[0] = 1
	fac[0] = 1
	for i := 1; i < n; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	num := make([]int, 1<<14)
	cnt := make([]int, 1<<14)
	for s := 0; s < 1<<n; s++ {
		for i := 1; i <= m; i++ {
			if s>>u[i]&1 != 0 && s>>v[i]&1 != 0 {
				num[s]++
			}
		}
		if s != 0 {
			cnt[s] = cnt[s^s&-s] + 1
		}
	}

	for s := 1; s < 1<<n; s++ {
		if s == (s & -s) {
			f[s] = 1
			continue
		}
		for t := (s - 1) & s; t > 0; t = (t - 1) & s {
			f[s] = (f[s] + f[s^t]*f[t]%mod*calc(s^t, t, num)%mod) % mod
		}
		f[s] = f[s] * invMod((cnt[s]-1)<<1) % mod
	}

	var dp [14][1 << 14]int
	for i := 0; i < n; i++ {
		for s := 0; s < 1<<n; s++ {
			if cnt[s]-1 == i {
				dp[i][s] = f[s]
			}
			if s == 0 {
				continue
			}
			x := s & -s
			k := s ^ x
			for t := (k - 1) & k; t != k; t = (t - 1) & k {
				if cnt[t] <= i {
					dp[i][s] = (dp[i][s] + dp[i-cnt[t]][s^t^x]*f[t^x]%mod) % mod
				}
			}
		}
	}

	p := 1
	for i := 1; i < n; i++ {
		p = p * m % mod
		fmt.Fprintln(out, dp[i][(1<<n)-1]*fac[i]%mod*invMod(p)%mod)
	}
}

func calc(s, t int, num []int) int {
	return num[s|t] - num[s] - num[t]
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

func invMod(a int) int {
	return powMod(a, mod-2)
}
