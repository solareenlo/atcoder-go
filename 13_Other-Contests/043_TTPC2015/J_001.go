package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = int(1e9 + 7)

var fac [1010]int
var inv [1010]int

func power(a, b int) int {
	if b == 0 {
		return 1
	}
	res := power(a*a%mod, b/2)
	if b%2 == 1 {
		res = (res * a) % mod
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, K int
	fmt.Fscan(reader, &N, &K)

	fac[0] = 1
	for i := 1; i <= N; i++ {
		fac[i] = fac[i-1] * i % mod
	}

	inv[N] = power(fac[N], mod-2)
	for i := N - 1; i >= 0; i-- {
		inv[i] = inv[i+1] * (i + 1) % mod
	}

	dp := make([]int, N+1)
	dp[0] = 1

	for k := K; k > 1; k-- {
		for i := N - k; i >= 0; i-- {
			cmb := int(1)
			for j := k; i+j <= N; j += k {
				cmb *= fac[k-1] * fac[N-i-j+k] % mod * inv[N-i-j] % mod * inv[k] % mod
				cmb %= mod
				dp[i+j] += dp[i] * cmb % mod * inv[j/k] % mod
				dp[i+j] %= mod
			}
		}
		dp[0] = 0
	}

	fmt.Println(dp[N])
}
