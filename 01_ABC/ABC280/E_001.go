package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	if n == 1 {
		fmt.Println(1)
		return
	}

	tmp := powMod(100, mod-2)
	var f [200005]int
	f[0] = 0
	f[1] = 1
	p1 := 100 - p
	for i := 2; i <= n; i++ {
		f[i] = ((p1*tmp%mod)*(f[i-1]+1)%mod + (p*tmp%mod)*(f[i-2]+1)%mod) % mod
	}
	fmt.Println(f[n])
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
