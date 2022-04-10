package main

import "fmt"

func main() {
	var k, n, m int
	fmt.Scan(&k, &n, &m)

	cnt := 0
	E := make([]int, 111)
	P := make([]int, 111)
	for i := 2; i*i <= m; i++ {
		if m%i != 0 {
			continue
		}
		cnt++
		P[cnt] = i
		for m%i == 0 {
			m /= i
			E[cnt]++
		}
	}
	if m > 1 {
		cnt++
		P[cnt] = m
		E[cnt] = 1
	}

	ans := 1
	for i := 1; i <= cnt; i++ {
		p := P[i]
		e := E[i]
		u := pow(p, e)
		x := n % u
		if x == 0 {
			tot := powMod(u, k)
			for d := 0; d < e; d++ {
				tot = (tot + mod - nCrMod(d+k-1, d)*powMod(p-1, k)%mod*powMod(p, k*e-k-d)%mod)
				tot %= mod
			}
			ans = ans * tot % mod
			continue
		}
		d := 0
		for x%p == 0 {
			x /= p
			d++
		}
		ans = ans * nCrMod(d+k-1, d) % mod * powMod(p-1, k-1) % mod * powMod(p, k*e-k-e+1) % mod
	}
	fmt.Println(ans)
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}

const mod = 998244353

func powMod(a, x int) int {
	a %= mod
	tot := 1
	for x != 0 {
		if (x & 1) != 0 {
			tot = tot * a % mod
		}
		a = a * a % mod
		x >>= 1
	}
	return tot
}

func nCrMod(m, n int) int {
	tot := 1
	for i := m; i >= m-n+1; i-- {
		tot = tot * (i % mod) % mod
	}
	for i := 1; i <= n; i++ {
		tot = tot * powMod(i, mod-2) % mod
	}
	return tot
}
