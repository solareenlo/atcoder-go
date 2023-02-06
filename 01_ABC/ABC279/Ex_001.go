package main

import "fmt"

const mod = 200003

var m, n int
var fac, ifac [mod + 10]int

func main() {
	fac[0], ifac[0], ifac[1] = 1, 1, 1
	for i := 1; i < mod; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	for i := 2; i < mod; i++ {
		ifac[i] = (mod - mod/i) * ifac[mod%i] % mod
	}
	for i := 1; i < mod; i++ {
		ifac[i] = ifac[i-1] * ifac[i] % mod
	}

	fmt.Scan(&n, &m)

	ans := 0
	for i := -1000000; i <= 1000000; i++ {
		j := i * (3*i + 1) / 2
		tmp := 1
		if i%2 != 0 {
			tmp = mod - 1
		}
		ans = (ans + lucas(m+n-j-1, 2*n-1)*tmp) % mod
	}
	fmt.Println(ans)
}

func c(n, m int) int {
	if n < m {
		return 0
	}
	return fac[n] * ifac[m] * ifac[n-m] % mod
}

func lucas(n, m int) int {
	if m == 0 {
		return 1
	}
	return c(n%mod, m%mod) * lucas(n/mod, m/mod) % mod
}
