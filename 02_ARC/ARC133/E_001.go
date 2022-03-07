package main

import "fmt"

func main() {
	var n, m, v, a int
	fmt.Scan(&n, &m, &v, &a)

	g := gcd(n, m)
	ans := 0
	for i := 1; i <= v; i++ {
		if i == 1 {
			ans = (ans + powMod(v, n+m)) % mod
			continue
		}
		pro := powMod((powMod(i-1, n/g)*powMod(v-i+1, m/g)+powMod(i-1, m/g)*powMod(v-i+1, n/g))%mod, g)
		if a >= i {
			ans = (ans + pro) % mod
		}
		ans = (ans + (powMod(v, n+m)+mod-pro)*(mod+1)/2) % mod
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
