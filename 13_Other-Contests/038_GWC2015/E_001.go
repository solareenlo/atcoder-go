package main

import "fmt"

func main() {
	var n, k, l int
	fmt.Scan(&n, &k, &l)

	var f [2000001]int
	f[0] = 1
	for i := 1; i <= 2000000; i++ {
		f[i] = f[i-1] * i % mod
	}

	a, K, N := 0, k, n
	if n == l {
		for i := 0; i < n; i++ {
			a = (a + powMod(K, gcd(n, i))) % mod
		}
		a = a * invMod(N) % mod
	} else {
		a = ((f[n+k-1] * invMod(f[n])) % mod) * invMod(f[k-1]) % mod
		if l%2 != 0 && n <= k {
			a = (a + ((f[k]*invMod(f[n]))%mod)*invMod(f[k-n])%mod) % mod
		}
	}
	fmt.Println(a)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

const mod = 1000000007

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
