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

	const K = 4000
	const MOD = 1000000007

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	inv := make([]int, K+2)
	inv[1] = 1
	for i := 2; i <= K+1; i++ {
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
	}
	fact := make([]int, K+2)
	finv := make([]int, K+2)
	fact[0] = 1
	finv[0] = 1
	for i := 1; i <= K+1; i++ {
		fact[i] = fact[i-1] * i % MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
	num := make([]int, K+1)
	num[0] = N % MOD
	for i := 1; i <= K; i++ {
		num[i] = num[i-1] * (N % MOD) % MOD
	}
	den := make([]int, K+1)
	for i := range den {
		den[i] = 1
	}
	for i := 0; i <= K; i++ {
		num[i] *= finv[i+1]
		num[i] %= MOD
		den[i] *= finv[i+1]
		den[i] %= MOD
	}
	S := make([]int, K+1)
	copy(S, num)
	for i := 1; i <= K; i++ {
		for j := 0; j < i; j++ {
			S[i] += MOD - S[j]*den[i-j]%MOD
			S[i] %= MOD
		}
	}
	for i := 1; i <= K; i++ {
		S[i] *= fact[i]
		S[i] %= MOD
	}
	F := make([]int, K+1)
	for i := 1; i <= K; i++ {
		F[i] = S[i] * inv[i] % MOD
		if i%2 == 0 {
			F[i] = (MOD - F[i]) % MOD
		}
	}
	f := make([]int, K+1)
	f[0] = 1
	for i := 1; i <= K; i++ {
		for j := 0; j <= i-1; j++ {
			f[i] += ((j + 1) * F[j+1] % MOD) * f[i-1-j] % MOD
		}
		f[i] %= MOD
		f[i] *= inv[i]
		f[i] %= MOD
	}
	for i := 0; i < Q; i++ {
		var a, d, k int
		fmt.Fscan(in, &a, &d, &k)
		pa := make([]int, k+1)
		pd := make([]int, k+1)
		pa[0] = 1
		pd[0] = 1
		for j := 1; j <= k; j++ {
			pa[j] = pa[j-1] * a % MOD
			pd[j] = pd[j-1] * d % MOD
		}
		ans := 0
		binom := 1
		for j := k; j >= 0; j-- {
			ans += f[j] * pa[k-j] % MOD * pd[j] % MOD * binom % MOD
			if j > 0 {
				binom *= N - j + 1
				binom %= MOD
				binom *= inv[k-j+1]
				binom %= MOD
			}
		}
		ans %= MOD
		fmt.Fprintln(out, ans)
	}
}
