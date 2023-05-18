package main

import "fmt"

func main() {
	const mod = 998244353

	var fac, finv, inv [200010]int
	fac[0] = 1
	fac[1] = 1
	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := 2; i < 200005; i++ {
		fac[i] = fac[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		finv[i] = finv[i-1] * inv[i] % mod
	}

	var N, M, K int
	fmt.Scan(&N, &M, &K)
	ans := 0
	for i := 1; i <= N; i++ {
		if M%i == 0 && K == 1 {
			ans += fac[N-1]
		} else if M%i > 0 && K > 1 {
			ans += fac[N-2]
		}
	}
	fmt.Println(ans % mod)
}
