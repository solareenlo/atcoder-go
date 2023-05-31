package main

import "fmt"

func main() {
	var N, M, P int
	fmt.Scan(&N, &M, &P)
	ans := 0
	for k := 1; k <= M; k++ {
		ans = (ans + powMod(2, k-1+N*(M-k))*((powMod(2, k*N-1)-powMod(2, (N+1)/2*k-1)+MOD)%MOD)%MOD) % MOD
	}
	fmt.Println(ans)
}

const MOD = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
