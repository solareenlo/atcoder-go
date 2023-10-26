package main

import "fmt"

func main() {
	const MOD = 1000000007
	const N = 10000005
	var inv, f [N]int
	var fac int
	var n int
	fmt.Scan(&n)
	n /= 2
	inv[1] = 1
	for i := 2; i <= n+1; i++ {
		inv[i] = inv[MOD%i] * (MOD - MOD/i) % MOD
	}
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] * (4*i - 2) % MOD * inv[i+1] % MOD
	}
	fac = f[n]
	for i := 1; i <= n+1; i++ {
		f[i-1] = f[i] * i % MOD
	}
	for i := n + 1; i >= 2; i-- {
		f[i] = f[i-2]
	}
	f[0] = 0
	f[1] = 0
	for i := 0; i < n; i++ {
		f[i+1] = (f[i+1] + f[i]*4) % MOD
	}

	var c string
	fmt.Scan(&c)
	c = " " + c
	s, as := 0, 0
	for i := n * 2; i >= 1; i-- {
		if c[i] == ')' {
			s++
		} else {
			as += s
		}
	}
	as %= MOD
	fc := 1
	for i := 1; i < n; i++ {
		fc = fc * i % MOD
	}
	as = (((as*f[n]+((n*n-as)%MOD)*((fac*n%MOD)*n%MOD-f[n]))%MOD)*fc%MOD)*fc%MOD + ((((((fac*fc%MOD)*n%MOD)*fc%MOD)*n%MOD)*n%MOD)*(n-1)%MOD)*inv[2]%MOD
	fmt.Println((as%MOD + MOD) % MOD)
}
