package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var inv, fact, ifact [1010]int
	var cnt, nxt, prv [1010]int

	var N int
	fmt.Fscan(in, &N)
	inv[1] = 1
	fact[0], fact[1] = 1, 1
	ifact[0], ifact[1] = 1, 1
	for i := 2; i <= N; i++ {
		inv[i] = (MOD - MOD/i) * inv[MOD%i] % MOD
		fact[i] = fact[i-1] * i % MOD
		ifact[i] = ifact[i-1] * inv[i] % MOD
	}

	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		cnt[x]++
	}

	prv[0] = 1
	for i := N; i >= 1; i-- {
		for j := 0; j <= N; j++ {
			nxt[j] = 0
		}
		for j := 0; j <= N-cnt[i]; j++ {
			if prv[j] != 0 {
				coeff1 := 1
				for u := 0; j+cnt[i] >= u*i; u++ {
					coeff2 := ((fact[j+cnt[i]] * coeff1 % MOD) * ifact[u] % MOD) * ifact[j+cnt[i]-u*i] % MOD
					nxt[j+cnt[i]-u*i] = (nxt[j+cnt[i]-u*i] + prv[j]*coeff2) % MOD
					coeff1 = coeff1 * ifact[i] % MOD
				}
			}
		}
		for j := 0; j < N+1; j++ {
			prv[j] = nxt[j]
		}
	}
	fmt.Println(prv[0])
}
